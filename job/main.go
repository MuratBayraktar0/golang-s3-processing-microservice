package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	// this job could be a lambda function triggered by s3

	ctx := context.TODO()

	// create an amazon s3 service client
	bucketBasics := NewBucketBasics(ctx)

	// initialize and Connect to MongoDB
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set in environment")
	}

	mongoClient, err := NewMongoDB(ctx, mongoURI)
	if err != nil {
		fmt.Println(err)
	}

	// defer disconnect from MongoDB
	defer mongoClient.client.Disconnect(ctx)

	handler(ctx, bucketBasics, mongoClient)

	fmt.Println("Job is done")
}

func handler(ctx context.Context, bucketBasics *BucketBasics, mongoClient *MongoDB) {
	list, err := bucketBasics.ListObjects("product-data")
	if err != nil {
		fmt.Println(err)
	}

	errChn := make(chan error, len(list))
	for _, item := range list {
		go func(list []types.Object, mongoClient *MongoDB) {
			lines, bytes, err := bucketBasics.GetS3JSONL("product-data", *item.Key)
			if err != nil {
				errChn <- err
				return
			}

			exist, err := mongoClient.isExist(ctx, len(lines), len(bytes), *item.Key)
			if err != nil {
				errChn <- err
				return
			}

			if exist {
				fmt.Printf("Already exist in the %s \n", *item.Key)
			} else {
				models := []mongo.WriteModel{}
				for _, line := range lines {
					if len(line) == 0 || line[0] != '{' {
						continue
					}

					product := Product{}
					err = json.Unmarshal(line, &product)
					if err != nil {
						fmt.Println(err)
					}

					models = append(models, mongo.NewInsertOneModel().SetDocument(product))
				}

				err := mongoClient.AddProduct(ctx, models)
				if err != nil {
					errChn <- err
					return
				}

				err = mongoClient.AddFile(ctx, len(lines), len(bytes), *item.Key)
				if err != nil {
					errChn <- err
					return
				}
			}
		}(list, mongoClient)
	}

	for i := 0; i < len(list); i++ {
		err := <-errChn
		fmt.Println(err)
	}
}
