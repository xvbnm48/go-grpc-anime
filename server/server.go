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
	"google.golang.org/grpc"
)

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

}

func main() {
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
