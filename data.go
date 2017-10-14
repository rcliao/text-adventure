package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
)

var salt = "CS-4660-fall-2017-" + strconv.Itoa(rand.Intn(460))
var states []State

// GenerateStateTree is the function to generate all mock data
func GenerateStateTree() {
	rand.Seed(201713374660)
	numOfNodes := 100
	minEdges := 3
	maxEdges := 17

	for i := 0; i < numOfNodes; i++ {
		currentState := createState(i)
		states = append(states, currentState)
	}

	for i := 0; i < numOfNodes; i++ {
		currentState := &states[i]
		for j := 0; j < minEdges+rand.Intn(maxEdges); j++ {
			randomNeighborIndex := rand.Intn(len(states))
			for randomNeighborIndex == i {
				randomNeighborIndex = rand.Intn(len(states))
			}
			randomNeighbor := states[randomNeighborIndex]
			randomNeighbor.Neighbors = []State{}
			currentState.Neighbors = append(currentState.Neighbors, randomNeighbor)
		}
	}
}

func createState(id int) State {
	locationName := locationNames[rand.Intn(len(locationNames))]
	currentLocation := NewLocation(
		locationName,
	)
	neighbors := []State{}
	return State{
		getMD5Hash(strconv.Itoa(id) + salt),
		*currentLocation,
		neighbors,
	}
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
