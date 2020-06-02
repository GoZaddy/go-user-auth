package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoDBClient is the mongodb client
var (
	mongoDBClient *mongo.Client
	err           error
	collection    *mongo.Collection
)

//Connect connects to MongoDB
func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	collection = mongoDBClient.Database("authPratice").Collection("users")

	err = mongoDBClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
}

//Disconnect disconnects from database
func Disconnect() {
	err := mongoDBClient.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
