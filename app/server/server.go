package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var queueName string = "go-sqs-queue"

type Server struct {
	SQS_URL    string
	SQS_Client *sqs.Client
}

type body struct {
	ID string 
	Message string `json:"message"`
}

func (server *Server) SetupServer() {
	http.HandleFunc("/", server.handler)
}

func (server *Server) handler(w http.ResponseWriter, r *http.Request) {
	// restrict to post request and use a struct to get info
	if r.Method != http.MethodPost {
		log.Println("not a valid request method")
		http.Error(w, "must be a post request", http.StatusBadRequest)
	}

	// pull the info from the body
	var b body
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println("cannot parse body")
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	if err != nil {
		log.Println("error getting sqsurl")
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	resp, err := server.SQS_Client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &server.SQS_URL,
		MessageBody: &b.Message,
	})
	if err != nil {
		log.Println("Got error when sending message")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	b.ID = *resp.MessageId
	log.Println("Sent Message: ")
	json.NewEncoder(w).Encode(b)
}
