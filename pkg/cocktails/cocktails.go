package cocktails

import (
	"context"
	"log"

	cocktailv1 "github.com/jestradaramos/sho/api/gen/go/v1"
)

// cocktailServiceServer implements the CocktailServiceServer API.
type CocktailServiceServer struct {
	cocktailv1.UnimplementedCocktailServiceServer
}

// CreateCocktail adds a cocktail
func (s *CocktailServiceServer) CreateCocktail(ctx context.Context, req *cocktailv1.CreateCocktailRequest) (*cocktailv1.CreateCocktailResponse, error) {
	name := req.Cocktail.Name
	log.Println("Got a request to create a", name)

	return &cocktailv1.CreateCocktailResponse{}, nil
}

// CreateCocktail adds a cocktail
func (s *CocktailServiceServer) GetCocktail(ctx context.Context, req *cocktailv1.GetCocktailRequest) (*cocktailv1.GetCocktailResponse, error) {
	name := req.Cocktail.Name
	log.Println("Got a request to get a", name)

	return &cocktailv1.GetCocktailResponse{}, nil
}

// CreateCocktail adds a cocktail
func (s *CocktailServiceServer) ListCocktail(ctx context.Context, req *cocktailv1.ListCocktailsRequest) (*cocktailv1.ListCocktailsResponse, error) {
	log.Println("Got a request to list all cocktails")

	return &cocktailv1.ListCocktailsResponse{}, nil
}

// CreateCocktail adds a cocktail
func (s *CocktailServiceServer) UpdateCocktail(ctx context.Context, req *cocktailv1.UpdateCocktailRequest) (*cocktailv1.UpdateCocktailResponse, error) {
	name := req.Cocktail.Name
	log.Println("Got a request to update", name)

	return &cocktailv1.UpdateCocktailResponse{}, nil
}

// CreateCocktail adds a cocktail
func (s *CocktailServiceServer) DeleteCocktail(ctx context.Context, req *cocktailv1.DeleteCocktailRequest) (*cocktailv1.DeleteCocktailResponse, error) {
	id := req.Id
	log.Println("Got a request to delete ", id)

	return &cocktailv1.DeleteCocktailResponse{}, nil
}
