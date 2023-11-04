package main

import "fmt"

// Creating a node
type Node struct {
	value int
	next  *Node
}

// print the linked list value
func printLinkedList(p *Node) {
	for p != nil {
		fmt.Printf("%d ", p.value)
		p = p.next
	}
}

func main() {
	// Initialize nodes
	var head *Node
	var one *Node = nil
	var two *Node = nil
	var three *Node = nil

	// Allocate memory
	one = &Node{}
	two = &Node{}
	three = &Node{}

	// Assign value values
	one.value = 1
	two.value = 2
	three.value = 3

	// Connect nodes
	one.next = two
	two.next = three
	three.next = nil

	// printing node-value
	head = one
	printLinkedList(head)
	printLinkedList(head)

}
