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
	"google.golang.org/grpc/internal/status"
)

var collection *mongo.Collection

type server struct{}
type animeItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (*server) CreateAnime(ctx context.Context, req *animepb.CreateAnimeRequest) (*animepb.CreateAnimeResponse, error) {
	fmt.Println("Create Anime!")
	anime := req.GetAnime()
	data := animeItem{
		Name:        anime.GetName(),
		Description: anime.GetDescription(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := collection.InsertOne(context.TODO(), data)
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
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &animepb.CreateAnimeResponse{
		Anime: &animepb.Anime{
			Id:          oid.Hex(),
			Name:        anime.GetName(),
			Description: anime.GetDescription(),
			CreatedAt:   anime.GetCreatedAt(),
			UpdatedAt:   anime.GetUpdatedAt(),
		},
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
