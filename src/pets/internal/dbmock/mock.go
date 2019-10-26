package dbmock

import (
	"context"

	"pets/internal/model"
)

type Mock struct {}

func (m *Mock) Find(context.Context) (model.Pets, error) {
	return model.Pets{
		model.Pet{
			Name: "bixa",
			Kind: "cat",
		},
	}, nil
}

func (m *Mock) InsertOne(context.Context, model.Pet) error {
	return nil
}

func (m *Mock) Ping(context.Context) error {
	return nil
}

func New() *Mock {
	return &Mock{}
}
