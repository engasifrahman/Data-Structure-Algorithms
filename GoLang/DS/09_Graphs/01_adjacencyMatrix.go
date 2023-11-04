package main

import "fmt"

type Graph struct {
	matrix [][]int
	size   int
}

func createGraph(size int) *Graph {
	adjMatrix := make([][]int, size)
	for i := 0; i < size; i++ {
		adjMatrix[i] = make([]int, size)
	}

	return &Graph{
		matrix: adjMatrix,
		size:   size,
	}
}

func (g *Graph) addEdge(v1, v2 int) {
	if v1 >= g.size || v2 >= g.size {
		fmt.Printf("Vertex %d and %d conbination not allowed due to graph size issue\n", v1, v2)
		return
	}

	g.matrix[v1][v2] = 1
	g.matrix[v2][v1] = 1
}

func (g *Graph) removeEdge(v1, v2 int) {
	if v1 >= g.size || v2 >= g.size || g.matrix[v1][v2] == 0 {
		fmt.Printf("No edge between %d and %d\n", v1, v2)
		return
	}

	g.matrix[v1][v2] = 0
	g.matrix[v2][v1] = 0
}

func (g *Graph) printMatrix() {
	for _, row := range g.matrix {
		for _, val := range row {
			fmt.Printf("%4d", val)
		}
		fmt.Println()
	}
}

func main() {
	g := createGraph(5)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(5, 4)

	g.removeEdge(5, 4)

	g.printMatrix()
}
