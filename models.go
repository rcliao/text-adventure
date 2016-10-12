package main

// State represent each state in the text-adventure game
type State struct {
	Location Location
	Hero     Hero
}

// Location present location with its name and description
type Location struct {
	Name        string
	Description string
}

// TODO: add a list of locations as enum

// NewLocation is a constructor pattern to generate location with random chance imported
func NewLocation(name, description string) *Location {
	return &Location{name, description}
}

// Hero is the player!
type Hero struct {
	HP int
}

// Event indicate what event happens to hero
type Event struct {
	Name        string
	Description string
	Effect      int
}

// NewEvent is a constructor for event
func NewEvent(name string, description string, effect int) *Event {
	return &Event{name, description, effect}
}

// TODO: generate a list of events
