package bun

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Cocktail struct {
	bun.BaseModel
	ID          uuid.UUID `bun:"type:uuid,default:gen_random_uuid(),pk"`
	Name        string    `bun:",notnull"`
	Notes       []string  `bun:",array"`
	Ingredients []string  `bun:",array"`
	Steps       []string  `bun:",array"`
}
