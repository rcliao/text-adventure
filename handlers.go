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

// HandleShowAllStates is internal state testing
func HandleShowAllStates() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(states)
	})
}

// HandleGetState return the state by id
func HandleGetState() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var t Action
		err := decoder.Decode(&t)
		if err != nil {
			panic("Oh no, right? OH NO PANIC!")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(findState(t.ID))
	})
}

// HandleStateTransition take state id and path to new location and return the new state
func HandleStateTransition() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var t Action
		err := decoder.Decode(&t)
		if err != nil {
			panic("Oh no, right? OH NO PANIC!")
		}
		var nextState State
		var nextID string
		for _, state := range states {
			if state.ID == t.ID {
				for _, n := range state.Neighbors {
					if n.ID == t.Action {
						nextID = n.ID
					}
				}
			}
		}
		if &nextID == nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		nextState = findState(nextID)

		if &nextState != nil {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Action{
				nextState.ID,
				nextState.Location.Name,
				nextState.Location.Event,
			})
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}

func findState(ID string) State {
	var result State
	for _, state := range states {
		if state.ID == ID {
			result = state
		}
	}
	return result
}
