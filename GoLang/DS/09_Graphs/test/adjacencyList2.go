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
		fmt.Printf("\nSorry, edge not allowed between %d and %d vertex because it's out of this graph scope!\n", v1, v2)
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
		fmt.Printf("\nSorry, edge not found between %d and %d vertex!\n", v1, v2)
		return false
	}

	g.removeDirectedEdge(v1, v2)
	g.removeDirectedEdge(v2, v1)

	return true
}

func (g *Graph) removeDirectedEdge(v1 int, v2 int) {
	if g.list[v1].data == v2 {
		g.list[v1] = g.list[v1].next
		return
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

func (g *Graph) printList() {
	for index, vertex := range g.list {
		fmt.Printf("\nVertex %d: ", index)

		for vertex != nil {
			fmt.Printf("→%2d", vertex.data)
			vertex = vertex.next
		}
	}
}

func main() {
	g := createGraph(5)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 1)
	g.addEdge(2, 1)
	g.addEdge(3, 1)
	g.addEdge(4, 1)
	g.addEdge(5, 1)

	g.printList()

	g.removeEdge(0, 2)
	g.removeEdge(0, 5)
	g.printList()
}