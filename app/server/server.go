package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type body struct {
	Id      string
	Message string `json:"message"`
}

func SetupServer() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
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

	b.Id = uuid.New().String()
	fmt.Println(b)
}
