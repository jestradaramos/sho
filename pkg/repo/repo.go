package repo

import (
	"context"

	cocktailv1 "github.com/jestradaramos/sho/api/gen/go/v1"
)

type Repo interface {
	CreateCocktail(context.Context, *cocktailv1.Cocktail) (string, error)
	// GetCocktail(string) cocktailv1.Cocktail
	// ListCocktail() []cocktailv1.Cocktail
	// UpdateCocktail(c cocktailv1.Cocktail) error
	// DeleteCocktail() error
}
