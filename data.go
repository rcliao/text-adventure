package main

import (
	"math/rand"

	"github.com/nu7hatch/gouuid"
)

var states []State

// GenerateStateTree is the function to generate all mock data
func GenerateStateTree() {
	rand.Seed(4660)
	hero := Hero{100}

	id := 0
	currentState := createState(hero)
	states = append(states, currentState)

	for i := 0; i < 3333; i++ {
		currentState := states[id]
		for i := 0; i < 3; i++ {
			var neighborState State
			neighborState = createState(currentState.Hero)
			states = append(states, neighborState)
			currentState.Neighbors = append(currentState.Neighbors, neighborState)
			states[id] = currentState
		}
		id++
	}
}

func createState(hero Hero) State {
	locationName := locationNames[rand.Intn(len(locationNames))]
	currentLocation := NewLocation(
		locationName,
	)
	neighbors := []State{}
	newHeroState := Hero{hero.HP}
	newHeroState.HP += currentLocation.Event.Effect
	u, _ := uuid.NewV4()
	return State{
		u.String(),
		*currentLocation,
		newHeroState,
		neighbors,
	}
}
