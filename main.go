package main

// TODO: store the state tree in memory
// TODO: add web-server to host the state tree
// TODO: add POST handler to handle state transition

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	// TODO: initialize the state tree with random rooms
	GenerateStateTree()
}

func main() {
	r := mux.NewRouter()

	r.Handle("/secret", Index()).Methods("GET")
	r.Handle("/secret/states", HandleShowAllStates()).Methods("GET")
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	r.Handle("/healthcheck", HandleHealthCheck()).Methods("GET")
	r.Handle("/getState", HandleGetState()).Methods("POST")
	r.Handle("/state", HandleStateTransition()).Methods("POST")

	log.Println("Server running at port 9000")

	err := http.ListenAndServe(":9000", r)
	log.Fatal(err)
}
