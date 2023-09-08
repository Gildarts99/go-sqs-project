package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) {
	fmt.Println("Body: ", sqsEvent.Records)
}

func main() {
	lambda.Start(handler)
}
