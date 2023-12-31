package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

const (
	uri      = "mongodb://localhost:27017"
	Database = "products-api"
)

type Collection string

const (
	ProductsCollection Collection = "products"
)

func GetMongoClient() (*mongo.Client, error) {

	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client

		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
