package main

import (
	"fmt"

	"github.com/anjanavitthal/graph-utils/pkg"
)

func main() {
	testGraph := &pkg.Graph{}
	v1 := testGraph.AddVertex(1, "first")
	v2 := testGraph.AddVertex(2, "two")
	v3 := testGraph.AddVertex(3, "three")
	v4 := testGraph.AddVertex(4, "four")
	v5 := testGraph.AddVertex(5, "five")
	v6 := testGraph.AddVertex(6, "six")
	v7 := testGraph.AddVertex(7, "seven")
	v8 := testGraph.AddVertex(8, "eight")

	testGraph.AddEdge(v1, v2)
	testGraph.AddEdge(v1, v3)
	testGraph.AddEdge(v2, v4)
	testGraph.AddEdge(v2, v5)
	testGraph.AddEdge(v3, v6)
	testGraph.AddEdge(v4, v7)
	testGraph.AddEdge(v5, v8)
	testGraph.AddEdge(v6, v8)
	testGraph.AddEdge(v8, v7)

	// testGraph.Print()
	foundVertex, err := testGraph.DepthFirstSearh(v2, 8)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Target Vertex Found with ID: %d and name %s\n", foundVertex.ID, foundVertex.Name)
	}
}
