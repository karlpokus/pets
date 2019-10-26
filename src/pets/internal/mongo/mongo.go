// package mongo wraps a mongo.Client and mongo.Collection
package mongo

import (
	"context"
	"fmt"
	"time"

	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pets/internal/model"
)

type Mongo struct {
	Client *mongo.Client
	Pets *mongo.Collection
}

func (m *Mongo) Find(ctx context.Context) (model.Pets, error) {
	var pets model.Pets
	cur, err := m.Pets.Find(ctx, bson.M{})
	if err != nil {
		return pets, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) { // cur.All(ctx, &list) is undefined. Bug submitted
		var pet model.Pet
		if err := cur.Decode(&pet); err != nil {
			return pets, err
		}
		pets = append(pets, pet) // TODO: paging
	}
	if err := cur.Err(); err != nil {
		return pets, err
	}
	return pets, nil
}

func (m *Mongo) InsertOne(ctx context.Context, pet model.Pet) error {
	_, err := m.Pets.InsertOne(ctx, pet)
	return err
}

func (m *Mongo) Ping(ctx context.Context) error {
	return m.Client.Ping(ctx, nil)
}

// Mongo implements the srv.Exiter interface
func (m *Mongo) Shutdown(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

// New returns a Mongo type
func New(name, host, port string) (*Mongo, error) {
	opts := options.Client()
	opts.SetAppName(name)
	opts.SetMonitor(apmmongo.CommandMonitor())
	opts.ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	m := &Mongo{}
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return m, fmt.Errorf("mongo connection err: %s", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return m, fmt.Errorf("mongo ping err: %s", err)
	}
	m.Client = client
	m.Pets = client.Database("pets").Collection("pets")
	return m, nil
}
