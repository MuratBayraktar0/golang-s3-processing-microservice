package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type BucketBasics struct {
	ctx      context.Context
	S3Client *s3.Client
}

func NewBucketBasics(ctx context.Context) *BucketBasics {
	// load the shared AWS configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-central-1"))
	if err != nil {
		fmt.Println(err)
	}

	// create an amazon s3 service client
	client := s3.NewFromConfig(cfg)

	return &BucketBasics{
		ctx:      ctx,
		S3Client: client,
	}
}

// File gets an object from a s3 bucket and stores it in a local file.
func (basics *BucketBasics) GetS3JSONL(bucketName string, objectKey string) ([][]byte, []byte, error) {
	result, err := basics.S3Client.GetObject(basics.ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
		return nil, nil, err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, nil, err
	}

	lines := bytes.Split(body, []byte("\n"))

	return lines, body, nil
}

// ListObjects lists the objects in a bucket.
func (basics *BucketBasics) ListObjects(bucketName string) ([]types.Object, error) {
	result, err := basics.S3Client.ListObjectsV2(basics.ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	var contents []types.Object
	if err != nil {
		log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", bucketName, err)
	} else {
		contents = result.Contents
	}
	return contents, err
}
