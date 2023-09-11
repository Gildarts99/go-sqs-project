package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-sqs-project/server"
)

func main() {
	// get the url of the sqs queue
	url := os.Getenv("SQS_URL")
	if url == "" {
		log.Println("error: SQS_URL environment variable is empty")
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	// setup an sqs client
	sqs_client := sqs.NewFromConfig(cfg)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	// create a server object
	server := server.Server{
		SQS_URL:    url,
		SQS_Client: sqs_client,
	}
	server.SetupServer()

	// make an http server
	// have a route for a post request that will allow for an individual to send a message
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
