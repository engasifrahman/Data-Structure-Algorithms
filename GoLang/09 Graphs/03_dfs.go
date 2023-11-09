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
	capacity int
	visited []int
	adjLists []*Vertex
}

func createGraph(capacity int) *Graph {
	g := &Graph{
		capacity: capacity,
	}
	g.visited = make([]int, capacity)
	g.adjLists = make([]*Vertex, capacity)

	for i := 0; i < capacity; i++ {
		g.visited[i] = 0
	}

	return g
}

func (g *Graph) addEdge(src, dest int) bool {

	if src < 0 || src >= g.capacity || dest < 0 || dest >= g.capacity {
		fmt.Printf("Sory, edge between %d and %d is not allowed due to graph capacity issue! \n", src, dest)
		return false
	}

	// Add edge from src to dest
	destVertex := createVertex(dest)
	destVertex.next = g.adjLists[src]
	g.adjLists[src] = destVertex

	// Add edge from dest to src
	srcVertex := createVertex(src)
	srcVertex.next = g.adjLists[dest]
	g.adjLists[dest] = srcVertex

	return true
}

func (g *Graph) printGraph() {
	for v := 0; v < g.capacity; v++ {
		current := g.adjLists[v]
		fmt.Printf("\nVertex %d: ", v)
		for current != nil {
			fmt.Printf("%d -> ", current.data)
			current = current.next
		}
	}

	fmt.Println()
}

func (g *Graph) dfs(start int) {
	adjList := g.adjLists[start]
	current := adjList

	g.visited[start] = 1
	fmt.Printf("Visited %d: ", start)
	fmt.Println(g.visited)

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
