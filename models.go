package main

import "math/rand"

// State represent each state in the text-adventure game
type State struct {
	ID        string   `json:"id"`
	Location  Location `json:"location"`
	Neighbors []State  `json:"neighbors"`
}

// Location present location with its name and description
type Location struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Name  string  `json:"name"`
	Event Event   `json:"-"`
}

// EventChance is to wrap the event with the chances
type EventChance struct {
	Event  Event
	Chance int
}

var locationDescriptionMap = map[string]string{
	"Dark Room":      "Neither light nor darkvision can penetrate the gloom in this chamber. An unnatural shade fills it, and the room's farthest reaches are barely visible. Near the room's center, you can just barely perceive a lump about the size of a human lying on the floor. (It might be a dead body, a pile of rags, or a sleeping monster that can take advantage of the room's darkness.)",
	"Room with cage": "A huge iron cage lies on its side in this room, and its gate rests open on the floor. A broken chain lies under the door, and the cage is on a rotting corpse that looks to be a hobgoblin. Another corpse lies a short distance away from the cage. It lacks a head.",
	"Hall Way":       "This short hall leads to another door. On either side of the hall, niches are set into the wall within which stand clay urns. One of the urns has been shattered, and its contents have spilled onto its shelf and the floor. Amid the ash it held, you see blackened chunks of something that might be bone.",
	"Chamber":        "This chamber was clearly smaller at one time, but something knocked down the wall that separated it from an adjacent room. Looking into that space, you see signs of another wall knocked over. It doesn't appear that anyone made an effort to clean up the rubble, but some paths through see more usage than others.",
	"Dire Tombs":     "This room is a tomb. Stone sarcophagi stand in five rows of three, each carved with the visage of a warrior lying in state. In their center, one sarcophagus stands taller than the rest. Held up by six squat pillars, its stone bears the carving of a beautiful woman who seems more asleep than dead. The carving of the warriors is skillful but seems perfunctory compared to the love a sculptor must have lavished upon the lifelike carving of the woman.",
	"Empty Room":     "You gaze into the room and hundreds of skulls gaze coldly back at you. They're set in niches in the walls in a checkerboard pattern, each skull bearing a half-melted candle on its head. The grinning bones stare vacantly into the room, which otherwise seems empty.",
}

var locationNames = []string{
	"Dark Room",
	"Room with cage",
	"Hall Way",
	"Chamber",
	"Dire Tombs",
	"Empty Room",
}
var locationEventsMap = map[string][]EventChance{
	"Dark Room": {
		EventChance{
			Event{
				"Loot",
				"You quicky loot the deadbody. Hey, there is still remaining healing potion!",
				20,
			},
			10,
		},
		EventChance{
			Event{
				"Darkness wispers",
				"You heard a wisper from the darkness. It drives you insane",
				-10,
			},
			30,
		},
		EventChance{
			Event{
				"Sleeping dwarve strike",
				"Turns out lump was a dwarf sleeping. He strikes and attacks you!",
				-30,
			},
			50,
		},
		EventChance{
			Event{
				"Stare in the darkness",
				"You stare into the darkness. Nothing happens",
				0,
			},
			100,
		},
	},
	"Room with cage": {
		EventChance{
			Event{
				"Treasure",
				"You found a healing potion in the cage!",
				20,
			},
			10,
		},
		EventChance{
			Event{
				"Hobgoblin attack!",
				"Besides the corpse, there is another hobgolbin hiding at the corner!",
				-20,
			},
			40,
		},
		EventChance{
			Event{
				"Nothing happens",
				"You look around the room. There is nothing",
				0,
			},
			100,
		},
	},
	"Hall Way": {
		EventChance{
			Event{
				"Shattered hallway",
				"Hall starts to shatter ... rocks are falling!",
				-30,
			},
			30,
		},
		EventChance{
			Event{
				"Skeleton warrior",
				"A skeleton warrior approaches you ... with a sword and shield.",
				-30,
			},
			40,
		},
		EventChance{
			Event{
				"Safe pass",
				"Nothing seems odd here. You pass the hall safely",
				0,
			},
			100,
		},
	},
	"Chamber": {
		EventChance{
			Event{
				"Orge smash!",
				"An orge smashes you with his hammer!",
				-30,
			},
			30,
		},
		EventChance{
			Event{
				"Centaurs!",
				"Centaurs stings you!",
				-50,
			},
			30,
		},
		EventChance{
			Event{
				"Idle",
				"Nothing happens",
				0,
			},
			100,
		},
	},
	"Dire Tombs": {
		EventChance{
			Event{
				"Spectre darins",
				"A spectre appears and start darining your soul.",
				-20,
			},
			30,
		},
		EventChance{
			Event{
				"Zombie attack",
				"At the corner, there is a zombie slowly approaches.",
				-20,
			},
			40,
		},
		EventChance{
			Event{
				"Idle",
				"Nohting happens",
				0,
			},
			100,
		},
	},
	"Empty Room": {
		EventChance{
			Event{
				"Idle",
				"Nohting happnes",
				0,
			},
			100,
		},
	},
}

// NewLocation is a constructor pattern to generate location with random chance imported
func NewLocation(name string) *Location {
	var event Event
	for i := 0; i < len(locationEventsMap[name]); i++ {
		if rand.Intn(100) < locationEventsMap[name][i].Chance {
			event = locationEventsMap[name][i].Event
			break
		}
	}
	return &Location{rand.Float64(), rand.Float64(), name, event}
}

// Hero is the player!
type Hero struct {
	HP int `json:"hp"`
}

// Event indicate what event happens to hero
type Event struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Effect      int    `json:"effect"`
}

// Action indicate a state trasition model
type Action struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Event  Event  `json:"event"`
}
