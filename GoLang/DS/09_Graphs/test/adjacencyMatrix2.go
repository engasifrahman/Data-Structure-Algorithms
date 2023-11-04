package main
import "fmt"

type Graph struct {
	matrix [][]int
	size int
}

func createGraph(size int) *Graph {
	adjMatrix := make([][]int, size)

	for i := 0; i < size; i++ {
		adjMatrix[i] = make([]int, size)
	}

	return &Graph{
		matrix: adjMatrix,
		size: size,
	}
}

func (g *Graph) addEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.size || v2 < 0 || v2 >= g.size {
		fmt.Printf("Sorry, edge not allowed between %d and %d vertex! \n", v1, v2)
		return false
	}

	g.matrix[v1][v2] = 1
	g.matrix[v2][v1] = 1

	return true
}

func (g *Graph) removeEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.size || v2 < 0 || v2 >= g.size {
		fmt.Printf("Sorry, edge not found between %d and %d vertex! \n", v1, v2)
		return false
	}

	g.matrix[v1][v2] = 0
	g.matrix[v2][v1] = 0

	return true
}

func (g *Graph) printGraph() {
	fmt.Println("Adjacency matrix:")
	for _, row := range g.matrix {
		for _, val := range row {
			fmt.Printf("%2d", val)
		}
		fmt.Println("")
	}
}


func main() {
	g := createGraph(5)
	g.addEdge(0, 1)
	g.addEdge(1, 1)
	g.addEdge(2, 1)
	g.addEdge(3, 1)
	g.addEdge(4, 1)
	g.addEdge(5, 1)

	g.printGraph()

	g.removeEdge(0, 1)
	g.removeEdge(5, 1)

	g.printGraph()
}