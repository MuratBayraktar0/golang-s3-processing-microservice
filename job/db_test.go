package main

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoDB_AddProduct(t *testing.T) {
	ctx := context.TODO()
	connectURI := "mongodb+srv://username:password@cluster0.os7e6fh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	db, err := NewMongoDB(ctx, connectURI)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	product1 := Product{
		ID:          1,
		Title:       "Product 1",
		Price:       100.0,
		Category:    "Category 1",
		Brand:       "Brand 1",
		URL:         "http://example.com/product1",
		Description: "Description 1",
	}
	product2 := Product{
		ID:          2,
		Title:       "Product 2",
		Price:       200.0,
		Category:    "Category 2",
		Brand:       "Brand 2",
		URL:         "http://example.com/product2",
		Description: "Description 2",
	}

	// Create test data
	model := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(product1),
		mongo.NewInsertOneModel().SetDocument(product2),
	}

	err = db.AddProduct(ctx, model)
	if err != nil {
		t.Fatalf("Failed to add products: %v", err)
	}

	// Verify the products are added
	collection := db.db.Collection("products")
	count, err := collection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		t.Fatalf("Failed to count products: %v", err)
	}

	expectedCount := int64(len(model))
	if count != expectedCount {
		t.Errorf("Expected %d products, got %d", expectedCount, count)
	}
}

func TestMongoDB_AddFile(t *testing.T) {
	ctx := context.TODO()
	connectURI := "mongodb+srv://username:password@cluster0.os7e6fh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	db, err := NewMongoDB(ctx, connectURI)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create test data
	productCount := 10
	bytes := 1024
	name := "test.txt"

	err = db.AddFile(ctx, productCount, bytes, name)
	if err != nil {
		t.Fatalf("Failed to add file: %v", err)
	}

	// Verify the file is added
	collection := db.db.Collection("files")
	filter := bson.D{
		{Key: "lines", Value: productCount},
		{Key: "bytes", Value: bytes},
		{Key: "name", Value: name},
	}
	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		t.Fatalf("Failed to find file: %v", result.Err())
	}
}

func TestMongoDB_isExist(t *testing.T) {
	ctx := context.TODO()
	connectURI := "mongodb+srv://username:password@cluster0.os7e6fh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	db, err := NewMongoDB(ctx, connectURI)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create test data
	productCount := 10
	bytes := 1024
	name := "test.txt"

	exist, err := db.isExist(ctx, productCount, bytes, name)
	if err != nil {
		t.Fatalf("Failed to check file existence: %v", err)
	}

	if !exist {
		t.Errorf("Expected file to exist, but it does not")
	}
}
