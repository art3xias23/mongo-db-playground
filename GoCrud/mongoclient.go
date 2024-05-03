package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbClient struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func GetMongoDbClient(dbPass string, dbUser string, ctx context.Context) (*MongoDbClient, error) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@172.28.224.1:27017/gocrud?authSource=gocrud"))

	if err != nil {
		fmt.Println("Error in mongodb client")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("connection established")

	database := client.Database("gocrud")
	collection := database.Collection("animals")

	return &MongoDbClient{
		Client:     client,
		Database:   database,
		Collection: collection,
	}, nil
}
