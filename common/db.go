package common

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// GetDbCollection returns a mongo.Collection from a db given a collection name
func GetDbCollection(collectionName string) *mongo.Collection {
  return db.Collection(collectionName)
}

// InitDb initializes the mongo database
func InitDb() error {
  uri := os.Getenv("MONGODB_URI")

  if uri == "" {
    return errors.New("MONGODB_URI is not set")
  }

  client, err := mongo.Connect(
    context.Background(),
    options.Client().ApplyURI(uri),
  )

  if err != nil {
    return err
  }

  dbName := os.Getenv("MONGODB_DATABASE")

  if dbName == "" {
    dbName = "go_demo"
  }

  db = client.Database(dbName)

  return nil
}

// CloseDb closes the mongo database
func CloseDb() error {
  return db.Client().Disconnect(context.Background())
}


