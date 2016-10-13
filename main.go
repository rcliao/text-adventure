package main

// TODO: store the state tree in memory
// TODO: add web-server to host the state tree
// TODO: add POST handler to handle state transition

import (
	"log"
	"math/rand"
	"net/http"
)

func init() {
	// TODO: initialize the state tree with random rooms
	GenerateStateTree()
}

func main() {
	rand.Seed(460)

	mux := http.NewServeMux()

	mux.Handle("/healthcheck", HandleHealthCheck())

	log.Println("Server running at port 9000")

	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}
