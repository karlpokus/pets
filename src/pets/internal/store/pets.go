package store

import (
	"context"

	"pets/internal/model"
)

type Pets interface {
	Find(context.Context) (model.Pets, error)
	InsertOne(context.Context, model.Pet) error
	Ping(context.Context) error
}
