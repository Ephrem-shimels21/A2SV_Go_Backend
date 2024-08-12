package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func ConnectDb() {
	log.Println("Connecting to MongoDB...")
	clientOpitions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOpitions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	TaskCollection = client.Database("TaskManager").Collection("tasks")
	UserCollection = client.Database("TaskManager").Collection("users")

	log.Println("Connected to MongoDB!")

}
