package repository

import (
	"context"
	"log"

	"product_service/internal/domain/entity"
	"product_service/internal/domain/interfaces"
	"product_service/internal/infrastructure/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	collection *mongo.Collection
}

func NewProductRepository(db *database.MongoDB, collectionName string) interfaces.ProductRepository {
	collection := db.DB.Collection(collectionName)
	return &Product{collection: collection}
}

func (r *Product) GetProducts(page, pageSize int64) ([]entity.ProductEntity, error) {
	ctx := context.Background()
	opts := options.Find()
	opts.SetSkip((page - 1) * pageSize)
	opts.SetLimit(pageSize)

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Printf("Failed to get products: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []entity.ProductEntity
	if err = cursor.All(ctx, &products); err != nil {
		log.Printf("Failed to decode products: %v", err)
		return nil, err
	}

	return products, nil
}

func (r *Product) GetProduct(id int64) (entity.ProductEntity, error) {
	ctx := context.Background()
	filter := bson.M{"_id": id}

	var product entity.ProductEntity
	if err := r.collection.FindOne(ctx, filter).Decode(&product); err != nil {
		log.Printf("Failed to get product: %v", err)
		return entity.ProductEntity{}, err
	}

	return product, nil
}
