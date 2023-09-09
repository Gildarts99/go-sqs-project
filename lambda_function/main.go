package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) {

	fmt.Println(sqsEvent.Records)
	// // for each record, ceaser cipher it and put it in s3 and a nobject with the path being the message id
	// for _, record := range sqsEvent.Records {
	// 	fmt.Printf("The record %s for event source %s = %s \n", record.MessageId, record.EventSource, record.Body)
	// }
}

func main() {
	lambda.Start(handler)
}
