package data

import (
	"context"
	"log"
	

	_"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var client *mongo.Client
var collection *mongo.Collection
var userCollection *mongo.Collection

func InitMongodb(mongodb_uri string)  {
	clientOptions := options.Client().ApplyURI(mongodb_uri)
	client,err := mongo.Connect(context.TODO(),clientOptions)

	if err != nil{
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

	collection = client.Database("taskManager").Collection("tasks")
	userCollection = client.Database("taskManager").Collection("users")
}

func DisconnectMongo() error {
	if client != nil{
		err := client.Disconnect(context.TODO())
		if err != nil {
			return err
		}
		client = nil
		collection = nil
		return nil
	}
	return nil
}