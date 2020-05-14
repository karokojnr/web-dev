package db

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDBCollection() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://karokojnr:karokojnr@cluster0-ubthk.gcp.mongodb.net"))
	if err != nil {
		return nil, err
	}
	collection := client.Database("vepago").Collection("users")
	return collection, nil
}
