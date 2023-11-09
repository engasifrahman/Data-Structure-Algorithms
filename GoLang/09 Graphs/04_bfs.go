package main

import "fmt"

type Vertex struct {
	data int
	next *Vertex
}

type Graph struct {
	capacity int
	visited  []int
	adjLists []*Vertex
}

type Queue struct {
	items []int
}

func createVertex(v int) *Vertex {
	return &Vertex{
		data: v,
		next: nil,
	}
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

func (q *Queue) enqeue(item int) bool {
	q.items = append(q.items, item)

	fmt.Printf("\n%d is enqueued, Queue: %d", item, q.items)

	return true
}

func (q *Queue) dequeue() (int, bool) {
	item := 0
	if q.isQueueEmpty() {
		fmt.Println("\nSorry, the queue is empty!")

		return item, false
	}

	item, q.items = q.items[0], q.items[1:]

	fmt.Printf("\n%d is dequeued, Queue: %d", item, q.items)

	return item, true
}

func (q *Queue) isQueueEmpty() bool {
	return len(q.items) == 0
}

func (g *Graph) bfs(start int) {
	var q Queue
	q.enqeue(start)
	g.visited[start] = 1

	for !q.isQueueEmpty() {
		vertex, _ := q.dequeue()
		fmt.Printf("\nVisited %d: ", vertex)
		fmt.Println(g.visited)

		current := g.adjLists[vertex]
		for current != nil {
			connectedVertex := current.data
			if g.visited[connectedVertex] == 0 {
				q.enqeue(connectedVertex)
				g.visited[connectedVertex] = 1
			}
			current = current.next
		}
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

	g.bfs(0)
}
