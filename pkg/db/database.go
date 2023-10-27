package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("error connecting to the database")
		return nil
	}
	DB = client
	return client
}

func CreateMongoCollection(dbName, collectionName string) (*mongo.Collection, error) {
	if DB == nil {
		return nil, fmt.Errorf("MongoDB client is not connected")
	}
	database := DB.Database(dbName)
	collection := database.Collection(collectionName)

	return collection, nil
}
