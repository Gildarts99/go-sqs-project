package client

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
	Bucket string
}

func NewS3Client() *S3Client {
	// get the S3_BUCKET and AWS_REGION from an environment variable
	bucket, exist := os.LookupEnv("S3_BUCKET")
	if !exist {
		log.Fatal("environment variable S3_BUCKET not set")
	}

	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	return &S3Client{
		Client: s3Client,
		Bucket: bucket,
	}
}
