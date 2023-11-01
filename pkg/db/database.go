package db

import (
	"context"
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

// func CreateMongoCollection() (*mongo.Collection, error) {
// 	dbName := "mydb"
// 	collectionName := "Todo"

// 	if DB == nil {
// 		return nil, fmt.Errorf("MongoDB client is not connected")
// 	}
// 	database := DB.Database(dbName)
// 	collection := database.Collection(collectionName)

// 	return collection, nil
// }
