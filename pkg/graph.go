package pkg

import (
	"errors"
	"fmt"

	"github.com/anjanavitthal/graph-utils/queue"
)

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	ID       int
	Name     string
	Adjacent []*Vertex
}

// add Vertex
func (g *Graph) AddVertex(id int, name string) *Vertex {

	if contains(g.Vertices, id) {
		err := fmt.Errorf("Vertex %v already exists in graph, not adding", id)
		fmt.Println(err.Error())
		return nil
	} else {
		// add vertex
		newVertex := &Vertex{ID: id, Name: name}

		g.Vertices = append(g.Vertices, newVertex)
		return newVertex
	}
}

func (g *Graph) AddEdge(fromVertex, toVertex *Vertex) {

	// fromVertex := g.getVertex(from)
	// toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v)", fromVertex.ID, toVertex.ID)
		fmt.Println(err.Error())
	} else if contains(fromVertex.Adjacent, toVertex.ID) {
		err := fmt.Errorf("edge already exists (%v --> %v)", fromVertex.ID, toVertex.ID)
		fmt.Println(err.Error())
	} else {
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
	}

}

func (g *Graph) getVertex(id int) *Vertex {

	for i, v := range g.Vertices {
		if v.ID == id {
			return g.Vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, id int) bool {

	for _, vertex := range s {
		if vertex.ID == id {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	// add vertex
	for _, vertex := range g.Vertices {
		fmt.Printf("Vertex %v : ", vertex.ID)
		fmt.Printf("[")
		for _, adjacent := range vertex.Adjacent {
			fmt.Printf(" %v ", adjacent.ID)
		}
		fmt.Println("]")
	}
}

func (g *Graph) BreadthFirstSearch(start *Vertex, target int) (*Vertex, error) {
	// Ensure the graph is not empty
	if len(g.Vertices) == 0 {
		return nil, errors.New("graph is empty")
	}

	if start == nil {
		start = g.Vertices[0]
	}

	found := g.getVertex(start.ID)
	if found == nil {
		return nil, errors.New("start vertex not found in the graph")
	}

	visited := make(map[int]bool)
	q := queue.NewQueue[*Vertex]()

	// Enqueue the start vertex and mark it as visited
	q.Enqueue(start)
	visited[start.ID] = true

	for !q.IsEmpty() {
		currentVertex, _ := q.Dequeue()
		fmt.Printf("Visited Vertex ID: %d\n", currentVertex.ID)

		if currentVertex.ID == target {
			return currentVertex, nil
		}

		for _, adjacent := range currentVertex.Adjacent {
			if !visited[adjacent.ID] {
				q.Enqueue(adjacent)
				visited[adjacent.ID] = true
			}
		}
	}
	return nil, errors.New("target vertex not found in the graph")
}

func (g *Graph) DepthFirstSearh(start *Vertex, target int) (*Vertex, error) {

	// Ensure the graph is not empty
	if len(g.Vertices) == 0 {
		return nil, errors.New("graph is empty")
	}

	// Use the first vertex as root if no start vertex is provided
	if start == nil {
		start = g.Vertices[0]
	}

	// To keep track of visited vertices and avoid infinite loops
	visited := make(map[int]bool)

	// Call the recursive DFS helper function
	return g.dfsRecursive(start, target, visited)
}

func (g *Graph) dfsRecursive(v *Vertex, target int, visited map[int]bool) (*Vertex, error) {
	// If the vertex is nil, return nil (though this shouldn't happen in a valid graph)
	if v == nil {
		return nil, nil
	}

	// Mark the current vertex as visited
	visited[v.ID] = true
	fmt.Printf("Visited Vertex ID: %d\n", v.ID)

	// If the target is found, return the vertex
	if v.ID == target {
		return v, nil
	}

	// Recursively visit each adjacent vertex
	for _, adjacent := range v.Adjacent {
		if !visited[adjacent.ID] {
			// Recursively search adjacent vertices
			if foundVertex, err := g.dfsRecursive(adjacent, target, visited); foundVertex != nil || err != nil {
				return foundVertex, err
			}
		}
	}

	// Return nil if the target vertex is not found
	return nil, nil
}
