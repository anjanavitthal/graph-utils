package main

import (
	"github.com/anjanavitthal/graph-utils/pkg"
)

func main() {
	testGraph := &pkg.Graph{}
	testGraph.AddVertex(1, "first")
	testGraph.AddVertex(2, "two")
	testGraph.AddVertex(3, "three")
	testGraph.AddVertex(4, "four")
	testGraph.AddVertex(4, "four")

	testGraph.AddEdge(1, 2)
	testGraph.AddEdge(1, 3)
	testGraph.AddEdge(1, 3)
	testGraph.AddEdge(1, 6)

	testGraph.Print()
}
