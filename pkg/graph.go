package pkg

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	id       int
	name     string
	adjacent []*Vertex
}

// add Vertex
func (g *Graph) AddVertex(id int, name string) {

	if contains(g.vertices, id) {
		err := fmt.Errorf("Vertex %v already exists in graph, not adding", id)
		fmt.Println(err.Error())
	} else {
		// add vertex
		g.vertices = append(g.vertices, &Vertex{id: id, name: name})
	}
}

func (g *Graph) AddEdge(from, to int) {

	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("edge already exists (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

func (g *Graph) getVertex(id int) *Vertex {

	for i, v := range g.vertices {
		if v.id == id {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, id int) bool {

	for _, vertex := range s {
		if vertex.id == id {
			return true
		}
	}
	return false
}
func (g *Graph) Print() {
	// add vertex
	for _, vertex := range g.vertices {
		fmt.Printf("Vertex %v : ", vertex.id)
		fmt.Printf("[")
		for _, adjacent := range vertex.adjacent {
			fmt.Printf(" %v ", adjacent.id)
		}
		fmt.Println("]")
	}
}
