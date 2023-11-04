package main

import "fmt"

type Vertex struct {
	data int
	next *Vertex
}

func createVertex(v int) *Vertex {
	return &Vertex{
		data: v,
		next: nil,
	}
}

type Graph struct {
	size int
	visited []int
	adjLists []*Vertex
}

func createGraph(size int) *Graph {
	g := &Graph{
		size: size,
	}
	g.visited = make([]int, size)
	g.adjLists = make([]*Vertex, size)

	for i := 0; i < size; i++ {
		g.visited[i] = 0
	}

	return g
}

func (g *Graph) addEdge(src, dest int) {
	// Add edge from src to dest
	destVertex := createVertex(dest)
	destVertex.next = g.adjLists[src]
	g.adjLists[src] = destVertex

	// Add edge from dest to src
	srcVertex := createVertex(src)
	srcVertex.next = g.adjLists[dest]
	g.adjLists[dest] = srcVertex
}

func (g *Graph) printGraph() {
	for v := 0; v < g.size; v++ {
		current := g.adjLists[v]
		fmt.Printf("\nVertex %d: ", v)
		for current != nil {
			fmt.Printf("%d -> ", current.data)
			current = current.next
		}
	}

	fmt.Println()
}

func (g *Graph) dfs(data int) {
	adjList := g.adjLists[data]
	current := adjList

	g.visited[data] = 1
	fmt.Printf("Visited %d: ", data)
	fmt.Println(g.visited)
	g.printGraph()


	for current != nil {
		connectedVertex := current.data

		if g.visited[connectedVertex] == 0 {
			g.dfs(connectedVertex)
		}
		current = current.next
	}
}

func main() {
	g := createGraph(4)
	fmt.Print(g)

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 3)

	g.printGraph()

	g.dfs(2)
}
