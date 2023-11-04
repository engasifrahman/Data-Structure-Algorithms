package main

import "fmt"

type Vertex struct {
	data int
	next *Vertex
}

type Graph struct {
	list []*Vertex
	size int
}

func createVertex(v int) *Vertex {
	return &Vertex{
		data: v,
		next: nil,
	}
}

func createGraph(size int) *Graph {
	adjList := make([]*Vertex, size)

	return &Graph{
		list: adjList,
		size: size,
	}
}

func (g *Graph) addEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.size || v2 < 0 || v2 >= g.size {
		fmt.Printf("\nSory, edge between %d and %d vertex is not allowed due to graph size issue! \n", v1, v2)
		return false
	}

	// Add edge from v1 to v2
	vertex2 := createVertex(v2)
	vertex2.next = g.list[v1]
	g.list[v1] = vertex2

	// Add edge from v2 to v1
	vertex1 := createVertex(v1)
	vertex1.next = g.list[v2]
	g.list[v2] = vertex1

	return true
}

func (g *Graph) removeEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.size || g.list[v1] == nil || v2 < 0 || v2 >= g.size || g.list[v2] == nil {
		fmt.Printf("\nSorry, no eadge found between %d and %d vertex! \n", v1, v2)

		return false
	}

	g.removeDirectedEdge(v1, v2)
	g.removeDirectedEdge(v2, v1)

	return true
}

func (g *Graph) removeDirectedEdge(v1 int, v2 int) {
	if g.list[v1].data == v2 {
		g.list[v1] = g.list[v1].next
	}

	current := g.list[v1]

	for current.next != nil {
		if current.next.data == v2 {
			current.next = current.next.next
			break
		}

		current = current.next
	}
}

func (g *Graph) printGraph() {
	for i := 0; i < g.size; i++ {
		fmt.Printf("\n Vertex %d: ", i)
		temp := g.list[i]

		for temp != nil {
			fmt.Printf(" → %d", temp.data)
			temp = temp.next
		}
		
	}
}

func main() {
	graph := createGraph(4)
	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(-1, -2)


	graph.printGraph()

	graph.removeEdge(0, 1)
	graph.removeEdge(-1, -1)

	graph.printGraph()
}
