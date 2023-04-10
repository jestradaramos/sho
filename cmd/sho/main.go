package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	cocktailv1 "github.com/jestradaramos/sho/api/gen/go/v1"
	"github.com/jestradaramos/sho/pkg/cocktails"
	"github.com/jestradaramos/sho/pkg/repo/bun"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Initialize DB
	repo, err := bun.New("postgres://postgres:example@localhost:5432/?sslmode=disable")
	if err != nil {
		return err
	}

	err = repo.SetUp(context.Background())
	if err != nil {
		return err
	}

	cocktailService := cocktails.New(repo)

	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	cocktailv1.RegisterCocktailServiceServer(server, cocktailService)
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
