package repo

import cocktailv1 "github.com/jestradaramos/sho/api/gen/go/v1"

type Repo interface {
	CreateCocktail(c cocktailv1.Cocktail) (id string)
	GetCocktail(id string) cocktailv1.Cocktail
	ListCocktail() []cocktailv1.Cocktail
	UpdateCocktail(c cocktailv1.Cocktail) error
	DeleteCocktail() error
}
