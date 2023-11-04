package main

import "fmt"

type Graph struct {
	matrix [][]int
	size int
}

func newGraph(size int) *Graph {
	adjMatrix := make([][]int, size)

	for i := 0; i < size; i++ {
		adjMatrix[i] = make([]int, size)
	}

	return &Graph {
		matrix: adjMatrix,
		size: size,
	}
}

func (g *Graph) addEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.size || v2 < 0 || v2 >= g.size {
		fmt.Printf("Sorry, edge between %d and %d vertex is not possible due to graph size issue! \n", v1, v2)
		return false
	}

	g.matrix[v1][v2] = 1
	g.matrix[v2][v1] = 1

	return true
}

func (g *Graph) removeEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.size || v2 < 0 || v2 >= g.size || g.matrix[v1][v2] == 0 {
		fmt.Printf("Sorry, no edge found between %d and %d! \n", v1, v2)
		return false
	}

	g.matrix[v1][v2] = 0
	g.matrix[v2][v1] = 0

	return true
}

func (g *Graph) print() {
	for _, row := range g.matrix {
		for _, val := range row {
			fmt.Printf("%4d", val)
		}

		fmt.Print("\n");
	}
}

func main() {
	graph := newGraph(5)

	graph.addEdge(0, 0)
	graph.addEdge(0, 3)
	graph.addEdge(1, 1)
	graph.addEdge(2, 2)
	graph.addEdge(3, 3)
	graph.addEdge(4, 4)
	graph.addEdge(5, 5)


	graph.removeEdge(0, 3)
	graph.removeEdge(5, 5)

	graph.print()
}