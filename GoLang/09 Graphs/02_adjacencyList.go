package main

import "fmt"

type Vertex struct {
	data int
	next *Vertex
}

type Graph struct {
	adjList  []*Vertex
	capacity int
}

func createGraph(capacity int) *Graph {
	return &Graph{
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

func (g *Graph) addEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.capacity || v2 < 0 || v2 >= g.capacity {
		fmt.Printf("Sory, edge between %d and %d is not allowed due to graph capacity issue! \n", v1, v2)
		return false
	}

	// Add edge from v1 to v2
	vertex2 := g.createVertex(v2)
	vertex2.next = g.adjList[v1]
	g.adjList[v1] = vertex2

	// Add edge from v2 to v1
	vertex1 := g.createVertex(v1)
	vertex1.next = g.adjList[v2]
	g.adjList[v2] = vertex1

	return true
}

func (g *Graph) removeEdge(v1 int, v2 int) bool {
	if v1 < 0 || v1 >= g.capacity || g.adjList[v1] == nil || v2 < 0 || v2 >= g.capacity || g.adjList[v2] == nil {
		fmt.Printf("Sory, no edge found between %d and %d! \n", v1, v2)

		return false
	}

	g.removeDirectedEdge(v1, v2)
	g.removeDirectedEdge(v2, v1)

	return true
}

func (g *Graph) removeDirectedEdge(v1 int, v2 int) {
	if g.adjList[v1].data == v2 {
		g.adjList[v1] = g.adjList[v1].next
	} else {
		current := g.adjList[v1]
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
	for v := 0; v < g.capacity; v++ {
		fmt.Printf("\n Vertex %d: ", v)
		temp := g.adjList[v]
		for temp != nil {
			fmt.Printf("%d -> ", temp.data)
			temp = temp.next
		}
		fmt.Printf("\n")
	}
}

func main() {
	graph := createGraph(4)
	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)

	graph.printGraph()

	fmt.Println("\nRemoving edge (0, 1)")
	graph.removeEdge(0, 1)

	graph.printGraph()
}
