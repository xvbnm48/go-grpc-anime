package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/xvbnm48/go-grpc-anime/animepb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection

type server struct{}
type animeItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	// CreatedAt   *timestamppb.Timestamp `bson:"createdAt"`
	// UpdatedAt   *timestamppb.Timestamp `bson:"updatedAt"`
}

func (*server) CreateAnime(ctx context.Context, req *animepb.CreateAnimeRequest) (*animepb.CreateAnimeResponse, error) {
	fmt.Println("Create Anime!")
	anime := req.GetAnime()
	data := animeItem{
		Name:        anime.GetName(),
		Description: anime.GetDescription(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID : %v", err),
		)
	}
	return &animepb.CreateAnimeResponse{
		Anime: &animepb.Anime{
			Id:          oid.Hex(),
			Name:        anime.GetName(),
			Description: anime.GetDescription(),
		},
	}, nil
}

func (*server) GetAnimeData(ctx context.Context, req *animepb.ReadAnimeRequest) (*animepb.ReadAnimeResponse, error) {
	fmt.Println("Read Anime!")
	animeId := req.GetAnimeId()
	oid, err := primitive.ObjectIDFromHex(animeId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID %v", err),
		)
	}
	// create an empty struct
	data := &animeItem{}
	res := collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(data)
	if res != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find anime with specified ID: %v", err),
		)
	}

	return &animepb.ReadAnimeResponse{
		Anime: &animepb.Anime{
			Id:          data.ID.Hex(),
			Name:        data.Name,
			Description: data.Description,
		},
	}, nil
}

func (*server) UpdateAnime(ctx context.Context, req *animepb.UpdateAnimeRequest) (*animepb.UpdateAnimeResponse, error) {
	fmt.Println("Update Anime!")
	anime := req.GetAnime()
	oid, err := primitive.ObjectIDFromHex(anime.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID %v", err),
		)
	}

	data := &animeItem{}
	res := collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(data)
	if res != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find anime with specified ID: %v", err),
		)
	}
	data.Name = anime.GetName()
	data.Description = anime.GetDescription()
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": data})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update anime: %v", err),
		)
	}

	return &animepb.UpdateAnimeResponse{
		Anime: &animepb.Anime{
			Id:          data.ID.Hex(),
			Name:        data.Name,
			Description: data.Description,
		},
	}, nil
}

func (*server) DeleteAnime(ctx context.Context, req *animepb.DeleteAnimeRequest) (*animepb.DeleteAnimeResponse, error) {
	fmt.Println("Delete Anime!")
	oid, err := primitive.ObjectIDFromHex(req.GetAnimeId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID %v", err),
		)
	}

	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete anime: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find anime with specified ID: %v", err),
		)
	}
	return &animepb.DeleteAnimeResponse{
		AnimeId: req.GetAnimeId(),
	}, nil
}

func main() {
	// if crash the program, print the error message
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("connecting to mongodb")
	var (
		client   *mongo.Client
		mongoURL = "mongodb://localhost:27017"
	)

	// connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("Database_anime").Collection("anime")

	fmt.Println("Hello, world!")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	reflection.Register(s)

	animepb.RegisterDatabaseAnimeServer(s, &server{})
	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// wait for ctrl-c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is received
	<-ch
	fmt.Println("Stopping server...")
	s.Stop()
	lis.Close()
}
