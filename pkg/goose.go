package goose

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ErrEmptyMongoURI = "mongoURI cannot be empty, please provide a valid mongoURI"
	ErrEmptyDatabase = "database cannot be empty, please provide a valid database"
)

type goose struct {
	Client   *mongo.Client
	database string
}

type GooseCollection struct {
	Collection *mongo.Collection
}

// NewGoose creates a new goose instance
func NewGoose(mongoURI, databse string) (*goose, error) {
	if mongoURI == "" {
		return nil, errors.New(ErrEmptyMongoURI)
	}

	if databse == "" {
		return nil, errors.New(ErrEmptyDatabase)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	return &goose{
		Client:   client,
		database: databse,
	}, nil

}

func (g *goose) GetCollection(collection string) *GooseCollection {
	mongoCollection := g.Client.Database(g.database).Collection(collection)
	return &GooseCollection{Collection: mongoCollection}
}
