package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func createNode(item int, next *Node) *Node {
	return &Node{
		data: item,
		next: next,
	}
}

func main() {
	var n *Node
	n = createNode(10, nil)

	fmt.Println(n)
	fmt.Printf("data = %d\n", n.data)
}