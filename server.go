package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhowliu/service"
)

type Response struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    [][]string `json:"data,omitempty"`
}

type Error struct {
	Errors string `json:"errors"`
}

type Body struct {
	Sentences []string `json:"sentences,omitempty"`
	Language  string   `json:"language"`
}

func Tokenize(res http.ResponseWriter, req *http.Request) {
	var body Body
	var tokens chan []string
	var results [][]string

	_ = json.NewDecoder(req.Body).Decode(&body)

	results = service.Tokenize(body.Sentences, body.Language, 10)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(Response{
		Success: true,
		Message: "OK",
		Data:    results,
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tokenize", Tokenize).Methods("POST")
	log.Printf("Sever is running on Port %d.\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", router))
}
