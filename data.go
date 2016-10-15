package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
)

var salt = "CS-4660-fall-2016-" + strconv.Itoa(rand.Intn(460))
var states []State

// GenerateStateTree is the function to generate all mock data
func GenerateStateTree() {
	rand.Seed(4660)
	hero := Hero{100}

	id := 0
	currentState := createState(hero, id)
	states = append(states, currentState)

	for i := 0; i < 3333; i++ {
		currentState := states[id]
		for i := 0; i < 3; i++ {
			var neighborState State
			neighborState = createState(currentState.Hero, len(states))
			states = append(states, neighborState)
			currentState.Neighbors = append(currentState.Neighbors, neighborState)
			states[id] = currentState
		}
		id++
	}
}

func createState(hero Hero, id int) State {
	locationName := locationNames[rand.Intn(len(locationNames))]
	currentLocation := NewLocation(
		locationName,
	)
	neighbors := []State{}
	newHeroState := Hero{hero.HP}
	newHeroState.HP += currentLocation.Event.Effect
	return State{
		getMD5Hash(strconv.Itoa(id) + salt),
		*currentLocation,
		newHeroState,
		neighbors,
	}
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
