package main

import (
	"log"
	"net/http"
	"os"
	"github.com/go-sqs-project/server"
)

func main() {
	// get the url of the sqs queue
	url := os.Getenv("SQS_URL")

	if url == "" {
		log.Println("error: SQS_URL environment variable is empty")
		os.Exit(1)
	}

	// setup the server routes
	server.SetupServer()

	// make an http server
	// have a route for a post request that will allow for an individual to send a message
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
