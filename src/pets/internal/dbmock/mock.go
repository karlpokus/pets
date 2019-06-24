package dbmock

import (
  "fmt"
	"context"

	"pets/internal/model"
)

type client struct {
	pets []model.Pet
	i    int
}

func (c *client) Database(s string) model.MongoClient {
	return c
}

func (c *client) Collection(s string) model.Collection {
	return c
}

func (c *client) Ping(ctx context.Context) error {
	return nil
}

func (c *client) Disconnect(ctx context.Context) error {
	return nil
}

func (c *client) Find(ctx context.Context, filter interface{}) (model.Cursor, error) {
	c.pets = []model.Pet{
		{"bixa", "cat"},
		{"rex", "cat"},
	}
	c.i = len(c.pets)
	return c, nil
}

func (c *client) InsertOne(ctx context.Context, v interface{}) error {
  _, ok := v.(model.Pet)
  if !ok {
    return fmt.Errorf("%v is not a Pet", v)
  }
  return nil
}

func (c *client) Close(ctx context.Context) error {
	return nil
}

func (c *client) Next(ctx context.Context) bool {
	c.i--
	return c.i >= 0
}

func (c *client) Decode(v interface{}) error {
	p, ok := v.(*model.Pet)
  if !ok {
    return fmt.Errorf("%v is not a Pet", v)
  }
  p.Name = c.pets[c.i].Name
	p.Kind = c.pets[c.i].Kind
	return nil
}

func (c *client) Err() error {
	return nil
}

func New() *client {
	return &client{}
}
