package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	cocktailv1 "github.com/jestradaramos/sho/api/gen/go/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	cocktailv1.RegisterCocktailServiceServer(server, &cocktailServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// cocktailServiceServer implements the CocktailServiceServer API.
type cocktailServiceServer struct {
	cocktailv1.UnimplementedCocktailServiceServer
}

// PutPet adds the pet associated with the given request into the PetStore.
func (s *cocktailServiceServer) CreateCocktail(ctx context.Context, req *cocktailv1.CreateCocktailRequest) (*cocktailv1.CreateCocktailResponse, error) {
	name := req.Cocktail.Name
	log.Println("Got a request to create a", name)

	return &cocktailv1.CreateCocktailResponse{}, nil
}
