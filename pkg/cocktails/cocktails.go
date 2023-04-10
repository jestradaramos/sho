package cocktails

import (
	"context"
	"fmt"
	"log"

	cocktailv1 "github.com/jestradaramos/sho/api/gen/go/v1"
	"github.com/jestradaramos/sho/pkg/repo"
	"github.com/jestradaramos/sho/pkg/repo/bun"
)

// cocktailServiceServer implements the CocktailServiceServer API.
type CocktailServiceServer struct {
	repo repo.Repo
}

func New(repo repo.Repo) *CocktailServiceServer {
	return &CocktailServiceServer{repo}
}

// CreateCocktail adds a cocktail and generates a UUID
func (s *CocktailServiceServer) CreateCocktail(ctx context.Context, req *cocktailv1.CreateCocktailRequest) (*cocktailv1.CreateCocktailResponse, error) {
	name := req.Cocktail.Name
	log.Println("Got a request to create a", name)
	id, err := s.repo.CreateCocktail(ctx, req.Cocktail)
	if err != nil {
		return nil, err
	}
	fmt.Println("New ID: ", id)

	return &cocktailv1.CreateCocktailResponse{}, nil
}

// GetCocktail retrieves a cocktail by it UUID
func (s *CocktailServiceServer) GetCocktail(ctx context.Context, req *cocktailv1.GetCocktailRequest) (*cocktailv1.GetCocktailResponse, error) {
	id := req.Id
	log.Println("Got a request to get a", id)
	cocktail, err := s.repo.GetCocktail(ctx, id)
	if err != nil {
		return nil, err
	}

	response := cocktailv1.Cocktail{
		Id:          cocktail.ID.String(),
		Name:        cocktail.Name,
		Notes:       cocktail.Notes,
		Ingredients: cocktail.Ingredients,
		Steps:       cocktail.Steps,
	}

	return &cocktailv1.GetCocktailResponse{Cocktail: &response}, nil
}

// ListCocktails Lists all available cocktails
func (s *CocktailServiceServer) ListCocktails(ctx context.Context, req *cocktailv1.ListCocktailsRequest) (*cocktailv1.ListCocktailsResponse, error) {
	log.Println("Got a request to list all cocktails")
	cocktails, err := s.repo.ListCocktail(ctx)
	if err != nil {
		return nil, err
	}

	protoCocktails := []*cocktailv1.Cocktail{}

	for _, c := range *cocktails {
		pc := fromRepoToProto(c)
		protoCocktails = append(protoCocktails, &pc)
	}

	return &cocktailv1.ListCocktailsResponse{Cocktails: protoCocktails}, nil
}

// UpdateCocktail updates an existing cocktail
func (s *CocktailServiceServer) UpdateCocktail(ctx context.Context, req *cocktailv1.UpdateCocktailRequest) (*cocktailv1.UpdateCocktailResponse, error) {
	name := req.Cocktail.Name
	log.Println("Got a request to update", name)
	err := s.repo.UpdateCocktail(ctx, req.Cocktail)
	if err != nil {
		return nil, err
	}

	return &cocktailv1.UpdateCocktailResponse{}, nil
}

// DeleteCocktail Removes a cocktail by its UUID
func (s *CocktailServiceServer) DeleteCocktail(ctx context.Context, req *cocktailv1.DeleteCocktailRequest) (*cocktailv1.DeleteCocktailResponse, error) {
	id := req.Id
	log.Println("Got a request to delete ", id)
	err := s.repo.DeleteCocktail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &cocktailv1.DeleteCocktailResponse{}, nil
}

func fromRepoToProto(cocktail bun.Cocktail) cocktailv1.Cocktail {
	p := cocktailv1.Cocktail{
		Id:          cocktail.ID.String(),
		Name:        cocktail.Name,
		Notes:       cocktail.Notes,
		Ingredients: cocktail.Ingredients,
		Steps:       cocktail.Steps,
	}

	return p
}
