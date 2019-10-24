// db returns a connection to a mongodb instance
package db

import (
	"context"
	"fmt"
	"time"

	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pets/internal/model"
)

type (
  Client struct {
  	c  *mongo.Client
  	db *mongo.Database
  }
  Collection struct {
  	c *mongo.Collection
  }
  Cursor struct {
  	c *mongo.Cursor
  }
)

func (c *Cursor) Close(ctx context.Context) error {
	return c.c.Close(ctx)
}

func (c *Cursor) Next(ctx context.Context) bool {
	return c.c.Next(ctx)
}

func (c *Cursor) Decode(v interface{}) error {
	return c.c.Decode(v)
}

func (c *Cursor) Err() error {
	return c.c.Err()
}

func (c *Collection) Find(ctx context.Context, v interface{}) (model.Cursor, error) {
	cur, err := c.c.Find(ctx, v)
	return &Cursor{cur}, err
}

func (c *Collection) InsertOne(ctx context.Context, v interface{}) error {
  _, err := c.c.InsertOne(ctx, v)
	return err
}

func (c *Client) Database(s string) model.MongoClient {
	c.db = c.c.Database(s)
	return c
}

func (c *Client) Collection(s string) model.Collection {
	return &Collection{c.db.Collection(s)}
}

func (c *Client) Ping(ctx context.Context) error {
	return c.c.Ping(ctx, nil)
}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.c.Disconnect(ctx)
}

// Client implements the srv.Exiter interface
func (c *Client) Shutdown(ctx context.Context) error {
	return c.Disconnect(ctx)
}

// New returns a validated client connection based on options read from env vars
func New(host, port string) (*Client, error) {
	clientOpts := options.Client().SetAppName("pets-service")
	clientOpts.SetMonitor(apmmongo.CommandMonitor())
	clientOpts.ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOpts)
	// connect err cannot be trusted!
	// is nil even though ping ctx deadline exceeded
	// tunnel was not started so there was no mongod running on the specified port
	// TODO: report bug
	err = client.Ping(ctx, nil)
	if err != nil {
		err = fmt.Errorf("mongodb connection err: %s", err)
	}
	return &Client{client, nil}, err
}
