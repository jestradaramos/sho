package bun

import (
	"context"
	"database/sql"

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

	return model.ID.String(), nil
}

func (c *CocktailRepo) GetCocktail(ctx context.Context, id string) (*Cocktail, error) {
	model := Cocktail{}
	// formatID, err := uuid.Parse(id)
	// if err != nil {
	// 	return nil, err
	// }
	err := c.DB.NewSelect().Model(&model).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func fromProtoToRepo(cocktail *cocktailv1.Cocktail) (Cocktail, error) {
	model := Cocktail{
		Name:        cocktail.Name,
		Notes:       cocktail.Notes,
		Ingredients: cocktail.Ingredients,
		Steps:       cocktail.Steps,
	}
	return model, nil
}