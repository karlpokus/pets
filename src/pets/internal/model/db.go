package model

import "context"

type (
	Cursor interface {
		Close(context.Context) error
		Next(context.Context) bool
		Decode(interface{}) error
		Err() error
	}

	Collection interface {
		Find(context.Context, interface{}) (Cursor, error)
		InsertOne(context.Context, interface{}) error
	}

	MongoClient interface {
		Database(string) MongoClient
		Collection(string) Collection
		Ping(context.Context) error
		Disconnect(context.Context) error
	}
)
