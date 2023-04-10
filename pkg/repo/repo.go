package repo

import (
	"context"

	cocktailv1 "github.com/jestradaramos/sho/api/gen/go/v1"
	"github.com/jestradaramos/sho/pkg/repo/bun"
)

type Repo interface {
	CreateCocktail(context.Context, *cocktailv1.Cocktail) (string, error)
	GetCocktail(context.Context, string) (*bun.Cocktail, error)
	ListCocktail(context.Context) (*[]bun.Cocktail, error)
	UpdateCocktail(context.Context, *cocktailv1.Cocktail) error
	DeleteCocktail(context.Context, string) error
}
