package main

// TODO: store the state tree in memory
// TODO: add web-server to host the state tree
// TODO: add POST handler to handle state transition

import (
	"log"
	"net/http"
)

func init() {
	// TODO: initialize the state tree with random rooms
	GenerateStateTree()
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/healthcheck", HandleHealthCheck())
	mux.Handle("/states", HandleShowAllStates())
	mux.Handle("/getState", HandleGetState())
	mux.Handle("/state", HandleStateTransition())

	log.Println("Server running at port 9000")

	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}
