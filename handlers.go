package main

import (
	"encoding/json"
	"net/http"
)

type healthCheck struct {
	// simply indicate if the service is up or down
	Alive bool `json:"alive"`
}

// HandleHealthCheck is simple healthcheck for testing if the http service is up
func HandleHealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&healthCheck{true})
	})
}
