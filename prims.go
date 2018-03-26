package main

import (
	"container/list"
	"log"
	"math"
)

// Source https://www.cse.ust.hk/~dekai/271/notes/L07/L07.pdf

//Edge is a connection bitween two points. Every edge has two terminal vertecies and a weight between the vertecies.
type Edge struct {
	source      string
	sink        string
	capacity    float32
	reverseEdge *Edge
}

// Graph is a complete graph with vertecies and edges between them.
type Graph struct {
	edges     []Edge
	vertecies map[string][]Edge
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(vertex string) {
	g.vertecies[vertex] = make([]Edge, 0)
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(source string, sink string, capacity float32) {
	edge := Edge{source, sink, capacity, nil}
	g.edges = append(g.edges, edge)
	g.vertecies[source] = append(g.vertecies[source], edge)
}

func (g Graph) getEdges(vertex string) ([]Edge, bool) {
	edges, ok := g.vertecies[vertex]
	return edges, ok
}

// Prims algorithm finds a minimal spanning tree for a graph. That simple
//
// Running time is O((|V| + |E|) times log(|V|))
func Prims(g Graph) []Edge {
	if len(g.edges) == 0 {
		return make([]Edge, 0)
	}

	r := g.edges[0] // pick any vertex as the root

	s := make([]Edge, 1)
	s = append(s, r)

	terminalEdges := list.New()
	terminalEdges.PushBack(r)

	hasAllNodes := true
	for hasAllNodes {
		var minimum float32 = math.MaxFloat32
		var smallestEdge Edge
		var smallestE *list.Element
		for e := terminalEdges.Front(); e != nil; e = e.Next() {
			terminalEdge := e.Value.(Edge)
			if terminalEdge.capacity < minimum {
				minimum = terminalEdge.capacity
				smallestEdge = terminalEdge
				smallestE = e
			}
		}

		terminalEdges.Remove(smallestE)
		edges, hasVertex := g.getEdges(smallestEdge.sink)
		if hasVertex {
			for _, edge := range edges {
				terminalEdges.PushBack(edge)
			}
		}

		s = append(s, smallestEdge)

		hasAllNodes = terminalEdges.Len() == 0
	}

	return s
}

// New creates a new graph and assigns its values
func New() Graph {
	edges := make([]Edge, 0)
	vertecies := make(map[string][]Edge)
	g := Graph{edges, vertecies}
	return g
}

func buildExampleGraph() Graph {
	var g = New()
	var vertecies = []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, vertex := range vertecies {
		g.AddVertex(vertex)
	}

	g.AddEdge("a", "b", 4)
	g.AddEdge("a", "c", 8)

	g.AddEdge("b", "a", 4)
	g.AddEdge("b", "c", 9)
	g.AddEdge("b", "d", 8)
	g.AddEdge("b", "e", 10)

	g.AddEdge("c", "a", 8)
	g.AddEdge("c", "b", 9)
	g.AddEdge("c", "d", 2)
	g.AddEdge("c", "f", 1)

	g.AddEdge("d", "b", 8)
	g.AddEdge("d", "c", 2)
	g.AddEdge("d", "e", 7)
	g.AddEdge("d", "f", 9)

	g.AddEdge("e", "b", 10)
	g.AddEdge("e", "d", 7)
	g.AddEdge("e", "f", 5)
	g.AddEdge("e", "g", 6)

	g.AddEdge("f", "c", 1)
	g.AddEdge("f", "d", 9)
	g.AddEdge("f", "e", 5)
	g.AddEdge("f", "g", 2)

	g.AddEdge("g", "e", 6)
	g.AddEdge("g", "f", 2)

	return g
}

func main() {
	g := buildExampleGraph()
	edges := Prims(g)
	for _, edge := range edges {
		log.Printf("Edge from %s to %s with cost %v", edge.source, edge.sink, edge.capacity)
	}
}
