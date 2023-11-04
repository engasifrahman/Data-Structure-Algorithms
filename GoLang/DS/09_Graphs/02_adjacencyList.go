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

func newVertex(v int) *Vertex {
	return &Vertex{
		data: v, 
		next: nil,
	}
}

func newGraph(size int) *Graph {
	adjList := make([]*Vertex, size)

	return &Graph{
		size: size,
		list: adjList,
	}
}

func (g *Graph) addEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.size || v2 < 0 || v2 >= g.size {
		fmt.Printf("Sory, edge between %d and %d is not allowed due to graph size issue! \n", v1, v2)
		return false
	}

	// Add edge from v1 to v2
	vertex2 := newVertex(v2)
	vertex2.next = g.list[v1]
	g.list[v1] = vertex2
	
	// Add edge from v2 to v1
	vertex1 := newVertex(v1)
	vertex1.next = g.list[v2]
	g.list[v2] = vertex1

	return true
}

func (g *Graph) removeEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.size || g.list[v1] == nil || v2 < 0 || v2 >= g.size || g.list[v2] == nil {
		fmt.Printf("Sory, no edge found between %d and %d! \n", v1, v2)

		return false
	}

	g.removeDirectedEdge(v1, v2)
	g.removeDirectedEdge(v2, v1)

	return true
}

func (g *Graph) removeDirectedEdge(v1 int, v2 int) {
	if g.list[v1].data == v2 {
		g.list[v1] = g.list[v1].next
	} else {
		current := g.list[v1]
		for current.next != nil {
			if current.next.data == v2 {
				current.next = current.next.next
				break
			}
			current = current.next
		}
	}
}

func (g *Graph) printGraph() {
	for v := 0; v < g.size; v++ {
		fmt.Printf("\n Vertex %d: ", v)
		temp := g.list[v]
		for temp != nil {
			fmt.Printf("%d -> ", temp.data)
			temp = temp.next
		}
		fmt.Printf("\n")
	}
}

func main() {
	graph := newGraph(4)
	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)

	graph.printGraph()

	fmt.Println("\nRemoving edge (0, 1)")
	graph.removeEdge(0, 1)

	graph.printGraph()
}
