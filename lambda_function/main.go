package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func caesar(r rune, shift int) rune {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) (events.APIGatewayProxyResponse, error) {
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

	// for each record, ceaser cipher it and put it in s3 and a nobject with the path being the message id
	for _, record := range sqsEvent.Records {
		ciphered := strings.Map(func(r rune) rune {
			return caesar(r, 1)
		}, record.Body)

		_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: &bucket,
			Key:    &record.MessageId,
			Body:   strings.NewReader(ciphered),
		})

		if err != nil {
			log.Fatalln("Error: ", err)
		}
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Success!\"",
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
