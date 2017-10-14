package main

import (
	"bytes"
	"container/heap"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	model "github.com/rcliao/text-adventure/models"
)

var fromState = flag.String("from", "", "from state id")
var toState = flag.String("to", "", "from state id")

var serverURL = "http://localhost:9000"
var stateAPI = serverURL + "/getState"
var transitionAPI = serverURL + "/state"

var emptyItem = Item{}

func main() {
	flag.Parse()

	if *fromState == "" || *toState == "" {
		flag.PrintDefaults()
		return
	}

	g := NewAdjacencyList()

	exploredRoom := []model.State{}
	openSet := []model.State{}
	state := getState(*fromState)

	openSet = append(openSet, state)
	counter := 0

	for len(openSet) > 0 {
		c := model.State{}
		c, openSet = openSet[0], openSet[1:]
		exploredRoom = append(exploredRoom, c)
		node := Node{Data{c.ID, c.Location.Name}}
		g.AddNode(node)
		fmt.Println("Exploring node", node, counter)
		counter++

		for _, n := range c.Neighbors {
			neighborState := getState(n.ID)
			action := getTransition(c.ID, n.ID)
			edge := Edge{
				FromNode: node,
				ToNode:   Node{Data{neighborState.ID, neighborState.Location.Name}},
				Weight:   action.Event.Effect,
			}
			g.AddEdge(edge)
			fmt.Println("Adding edge", edge)
			if checkExist(exploredRoom, neighborState) || checkExist(openSet, neighborState) {
				continue
			}
			openSet = append(openSet, neighborState)
		}
	}

	fromNode := Node{Data{*fromState, "Empty Room"}}
	toNode := Node{Data{*toState, "Dark Room"}}
	bfs := search(func(g Graph, n1, n2 Node) int { return 1 })
	dijkstra := search(func(g Graph, n1, n2 Node) int { return g.Distance(n1, n2) * -1 })

	fmt.Println("BFS:\n", prettyPrintEdges(bfs(g, fromNode, toNode)))
	fmt.Println("Dijkstra:\n", prettyPrintEdges(dijkstra(g, fromNode, toNode)))
}

func checkExist(rooms []model.State, room model.State) bool {
	for _, r := range rooms {
		if r.ID == room.ID {
			return true
		}
	}
	return false
}

func getState(id string) model.State {
	values := model.State{
		ID: id,
	}
	bodyString, _ := json.Marshal(values)
	resp, err := request(stateAPI, bytes.NewBuffer(bodyString))
	if err != nil {
		log.Fatal("has error reaching server", err)
	}
	var data model.State
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Failed to parse response as JSON", err)
	}
	return data
}

func getTransition(id, action string) model.Action {
	values := model.Action{
		ID:     id,
		Action: action,
	}
	bodyString, _ := json.Marshal(values)
	resp, err := request(transitionAPI, bytes.NewBuffer(bodyString))
	if err != nil {
		log.Fatal("has error reaching server", err)
	}
	var data model.Action
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Failed to parse response as JSON", err)
	}
	return data
}

func request(url string, body *bytes.Buffer) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal("Failed to make request to server", err)
	}

	req.Header.Add("User-Agent", "Solution-App")
	req.Header.Add("Accept", "application/json")

	return client.Do(req)
}

// Node define generic node to hold data inside
type Node struct {
	Data Data
}

// Data is DTO (Data Transfer Object) wrapping the state and room name
type Data struct {
	ID   string
	Name string
}

// Edge defines the edge between node for the graph
type Edge struct {
	FromNode Node
	ToNode   Node
	Weight   int
}

// Graph is a generic interface for graph
type Graph interface {
	Len() int
	Adjacent(n1, n2 Node) bool
	Neighbors(node Node) []Node
	AddNode(node Node) bool
	RemoveNode(node Node) bool
	AddEdge(edge Edge) bool
	RemoveEdge(edge Edge) bool
	Distance(n1, n2 Node) int
}

// Search defines the common search API
type Search func(graph Graph, fromNode Node, toNode Node) []Edge

func search(getDistance func(g Graph, n1, n2 Node) int) Search {
	return func(graph Graph, fromNode Node, toNode Node) []Edge {
		edges := []Edge{}
		// start with one item inside
		frontier := make(PriorityQueue, 1)
		parents := make(map[Node]Node)
		distances := make(map[Node]float64)
		explored := []Node{}

		frontier[0] = &Item{
			value:    fromNode,
			priority: 0,
		}
		heap.Init(&frontier)

		parents[fromNode] = Node{}
		distances[fromNode] = 0

		for frontier.Len() > 0 {
			current := heap.Pop(&frontier).(*Item).value
			explored = append(explored, current)

			if current == toNode {
				// found solution
				return backTrack(graph, parents, current)
			}

			for _, n := range graph.Neighbors(current) {
				newCost := distances[current] + float64(getDistance(graph, current, n))
				if _, okay := distances[n]; !okay || newCost < distances[n] {
					priority := newCost
					if containsNode(explored, n) {
						continue
					}
					item := &Item{
						value:    n,
						priority: priority,
					}
					if existingItem := frontier.Contains(n); *existingItem != emptyItem {
						frontier.update(existingItem, item.value, priority)
					} else {
						heap.Push(&frontier, item)
					}
					distances[n] = newCost
					parents[n] = current
				}
			}
		}

		return edges
	}
}

func containsNode(nodes []Node, node Node) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}

	return false
}

func backTrack(graph Graph, parents map[Node]Node, current Node) []Edge {
	edges := []Edge{}
	c := current
	emptyNode := Node{}

	for c != emptyNode {
		parent := parents[c]
		edge := Edge{parent, c, graph.Distance(parent, c)}
		edges = append(edges, edge)
		c = parent
	}
	// remove last item
	edges = edges[:len(edges)-1]
	reverse(edges)

	return edges
}

func reverse(ss []Edge) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func prettyPrintEdges(edges []Edge) string {
	result := ""
	hp := 0

	for _, edge := range edges {
		result += edge.FromNode.Data.Name + "(" + edge.FromNode.Data.ID + ")" +
			":" + edge.ToNode.Data.Name + "(" + edge.ToNode.Data.ID + ")" +
			":" + strconv.Itoa(edge.Weight) + "\n"
		hp += edge.Weight
	}

	result += "hp:" + strconv.Itoa(hp)

	return result
}
