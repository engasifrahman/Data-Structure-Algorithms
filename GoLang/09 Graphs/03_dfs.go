package main

import "fmt"

type Vertex struct {
	data int
	next *Vertex
}

type Graph struct {
	visited  []int
	adjList  []*Vertex
	capacity int
}

func createGraph(capacity int) *Graph {
	return &Graph{
		visited:  make([]int, capacity),
		adjList:  make([]*Vertex, capacity),
		capacity: capacity,
	}
}

func (g *Graph) createVertex(v int) *Vertex {
	return &Vertex{
		data: v,
		next: nil,
	}
}

func (g *Graph) addEdge(src, dest int) bool {

	if src < 0 || src >= g.capacity || dest < 0 || dest >= g.capacity {
		fmt.Printf("Sory, edge between %d and %d is not allowed due to graph capacity issue! \n", src, dest)
		return false
	}

	// Add edge from src to dest
	destVertex := g.createVertex(dest)
	destVertex.next = g.adjList[src]
	g.adjList[src] = destVertex

	// Add edge from dest to src
	srcVertex := g.createVertex(src)
	srcVertex.next = g.adjList[dest]
	g.adjList[dest] = srcVertex

	return true
}

func (g *Graph) printGraph() {
	for v := 0; v < g.capacity; v++ {
		current := g.adjList[v]
		fmt.Printf("\nVertex %d: ", v)
		for current != nil {
			fmt.Printf("%d -> ", current.data)
			current = current.next
		}
	}

	fmt.Println()
}

func (g *Graph) dfs(start int) {
	g.visited[start] = 1

	fmt.Printf("\nVisited %d: %v", start, g.visited)

	current := g.adjList[start]
	for current != nil {
		connectedVertex := current.data
		if g.visited[connectedVertex] == 0 {
			g.dfs(connectedVertex)
		}

		current = current.next
	}
}

func main() {
	g := createGraph(5)
	fmt.Print(g)

	g.addEdge(0, 3)
	g.addEdge(0, 2)
	g.addEdge(0, 1)
	g.addEdge(1, 2)
	g.addEdge(2, 4)

	g.printGraph()

	g.dfs(0)
}
