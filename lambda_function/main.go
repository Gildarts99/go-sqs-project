package main

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/go-sqs-lambda/client"
	"github.com/go-sqs-lambda/crypto"
)

// create a new s3 client
var s3_client = client.NewS3Client()

// create a new ssm client
var ssm_client = client.NewSSMClient()

func handler(ctx context.Context, sqsEvent events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	secret, err := ssm_client.GetParameter("/go-sqs-project/ENCRYPTION_KEY", true)
	if err != nil {
		panic(err)
	}

	// for each record, ceaser cipher it and put it in s3 and a nobject with the path being the message id
	for _, record := range sqsEvent.Records {
		encrypted := crypto.Encrypt(record.Body, secret)

		_, err := s3_client.Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: &s3_client.Bucket,
			Key:    &record.MessageId,
			Body:   strings.NewReader(string(encrypted)),
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
