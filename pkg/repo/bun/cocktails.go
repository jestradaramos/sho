package bun

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	cocktailv1 "github.com/jestradaramos/sho/api/gen/go/v1"
)

type CocktailRepo struct {
	*bun.DB
}

func New(dsn string) (*CocktailRepo, error) {
	// dsn = "postgres://postgres:@localhost:5432/test?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return &CocktailRepo{db}, nil
}

func (c *CocktailRepo) SetUp(ctx context.Context) error {
	_, err := c.DB.NewCreateTable().Model((*Cocktail)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *CocktailRepo) CreateCocktail(ctx context.Context, cocktail *cocktailv1.Cocktail) (string, error) {
	model, err := fromProtoToRepo(cocktail)
	if err != nil {
		return "", err
	}
	_, err = c.DB.NewInsert().Model(&model).Exec(ctx)
	if err != nil {
		return "", err
	}

	return model.ID, nil
}

func fromProtoToRepo(cocktail *cocktailv1.Cocktail) (Cocktail, error) {
	ingredients := []string{}
	for _, i := range cocktail.Ingredients {
		in := fmt.Sprintf("%voz %v", i.Amount, i.Item)
		ingredients = append(ingredients, in)
	}
	model := Cocktail{
		Name:        cocktail.Name,
		Notes:       cocktail.Notes,
		Ingredients: ingredients,
		Steps:       cocktail.Steps,
	}
	return model, nil
}
