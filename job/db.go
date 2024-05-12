package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDB(ctx context.Context, connectURI string) (*MongoDB, error) {
	// Connect to MongoDB
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://username:password@cluster0.os7e6fh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		client: mongoClient,
		db:     mongoClient.Database("product"),
	}, nil
}

func (mongodb *MongoDB) isExist(ctx context.Context, productCount, bytes int, name string) (bool, error) {
	collection := mongodb.db.Collection("files")

	filter := bson.D{
		{Key: "lines", Value: productCount},
		{Key: "bytes", Value: bytes},
		{Key: "name", Value: name},
	}

	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil
}

func (mongodb *MongoDB) AddProduct(ctx context.Context, model []mongo.WriteModel) error {
	collection := mongodb.db.Collection("products")

	ordered := false
	_, err := collection.BulkWrite(ctx, model, &options.BulkWriteOptions{
		Ordered: &ordered, // If false, execution will continue after one fails
	})
	if err != nil {
		return err
	}

	return nil
}

func (mongodb *MongoDB) AddFile(ctx context.Context, productCount, bytes int, name string) error {
	collection := mongodb.db.Collection("files")

	file := bson.D{
		{Key: "lines", Value: productCount},
		{Key: "bytes", Value: bytes},
		{Key: "name", Value: name},
	}
	_, err := collection.InsertOne(ctx, file)
	if err != nil {
		return err
	}
	return nil
}
