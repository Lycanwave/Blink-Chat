package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"         // Use v2 for mongo
	"go.mongodb.org/mongo-driver/v2/mongo/options" // Use v2 for options
)

var Client *mongo.Client
var Context context.Context

func ConnectToMongo(uri string) error {

	// set client options
	clientOptions := options.Client().ApplyURI(uri)

	Context, _ = context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return fmt.Errorf("could not connect to MongoDB %v", err)
	}

	Client = client

	if err := Client.Ping(Context, nil); err != nil {
		return fmt.Errorf("could not ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")
	return nil

}

func GetCollection(database, collection string) *mongo.Collection {
	return Client.Database(database).Collection(collection)
}

func Disconnect() {
	if err := Client.Disconnect(Context); err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB")
}
