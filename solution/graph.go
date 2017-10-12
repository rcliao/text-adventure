package main

// AdjacencyList is a graph representation uses adjacency list
type AdjacencyList struct {
	data map[Node][]Edge
}

// NewAdjacencyList is a constructor pattern for adjacency list graph
func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{
		data: make(map[Node][]Edge),
	}
}

// Len return number of nodes in graph
func (g *AdjacencyList) Len() int { return len(g.data) }

// Adjacent checks if n1 and n2 are adjacent(connected)
func (g *AdjacencyList) Adjacent(n1 Node, n2 Node) bool {
	if !g.checkNodeExist(n1) || !g.checkNodeExist(n2) {
		return false
	}
	edges := g.data[n1]
	for _, edge := range edges {
		if edge.ToNode == n2 {
			return true
		}
	}

	return false
}

// Neighbors returns all the nodes that is adjacent from node
func (g *AdjacencyList) Neighbors(node Node) []Node {
	if !g.checkNodeExist(node) {
		return []Node{}
	}
	return getToNodesFromEdges(g.data[node])
}

// AddNode adds a new node to graph, return false if node already existed
func (g *AdjacencyList) AddNode(node Node) bool {
	if g.checkNodeExist(node) {
		return false
	}
	g.data[node] = []Edge{}
	return true
}

// RemoveNode removes an existing node from graph, return false if node does not exist
func (g *AdjacencyList) RemoveNode(node Node) bool {
	if !g.checkNodeExist(node) {
		return false
	}
	delete(g.data, node)
	return true
}

// AddEdge adds a new edge to graph, returns false if edge already existed
func (g *AdjacencyList) AddEdge(edge Edge) bool {
	node := edge.FromNode
	if !g.checkNodeExist(node) {
		// if the from node doesn't exist, add the node into dictionary
		g.data[node] = []Edge{}
	}
	if !g.checkNodeExist(edge.ToNode) {
		// if the from node doesn't exist, add the node into dictionary
		g.data[edge.ToNode] = []Edge{}
	}
	if checkEdgeInSlice(g.data[node], edge) {
		return false
	}
	g.data[node] = append(g.data[node], edge)
	return true
}

// RemoveEdge removes an existing edge from graph, returns false if edge doesn't exist
func (g *AdjacencyList) RemoveEdge(edge Edge) bool {
	node := edge.FromNode
	if !g.checkNodeExist(node) {
		return false
	}
	if !checkEdgeInSlice(g.data[node], edge) {
		return false
	}
	g.data[node] = removeEdge(g.data[node], edge)
	return true
}

func (g *AdjacencyList) Distance(n1, n2 Node) int {
	if !g.checkNodeExist(n1) || !g.checkNodeExist(n2) {
		return 0
	}
	for _, e := range g.data[n1] {
		if e.ToNode == n2 {
			return e.Weight
		}
	}
	return 0
}

func (g *AdjacencyList) checkNodeExist(node Node) bool {
	_, okay := g.data[node]
	return okay
}

func getToNodesFromEdges(edges []Edge) []Node {
	nodes := []Node{}
	for _, edge := range edges {
		nodes = append(nodes, edge.ToNode)
	}

	return nodes
}

func removeEdge(edges []Edge, edge Edge) []Edge {
	for i, e := range edges {
		if e == edge {
			return append(edges[:i], edges[i+1:]...)
		}
	}
	return edges
}

func checkEdgeInSlice(edges []Edge, edge Edge) bool {
	for _, e := range edges {
		if e == edge {
			return true
		}
	}
	return false
}
