package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) {
	fmt.Println("Body: ", event.Body)
}

func main() {
	lambda.Start(handler)
}
