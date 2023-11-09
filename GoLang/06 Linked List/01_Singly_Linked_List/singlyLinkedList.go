package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert at the beginning of the linked list
func (list *LinkedList) insertAtBeginning(data int) {
	newNode := &Node{
		data: data, 
		next: list.head,
	}

	list.head = newNode
}

// Insert at the end of the linked list
func (list *LinkedList) insertAtEnd(data int) {
	newNode := &Node{
		data: data, 
		next: nil,
	}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Insert after a given node in the linked list
func (list *LinkedList) insertAfter(node *Node, data int) {
	if node == nil {
		fmt.Println("Invalid node provided")
		return
	}

	newNode := &Node{
		data: data, 
		next: node.next,
	}

	node.next = newNode
}

// Delete a node from the linked list
func (list *LinkedList) deleteNode(data int) bool {
	if list.head == nil {
		return false
	}

	// Case 1: If the node to be deleted is the head node
	if list.head.data == data {
		list.head = list.head.next
		return true
	}

	// Case 2: If the node to be deleted is not the head node
	deleted := false
	current := list.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			deleted = true
			break
		}

		current = current.next
	}

	return deleted
}

// Traverse the linked list
func (list *LinkedList) traverse() {
	current := list.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.insertAtEnd(10)
	list.insertAtBeginning(20)
	list.insertAtBeginning(30)
	list.insertAtEnd(40)
	list.insertAfter(list.head.next, 50)

	list.traverse()

	list.deleteNode(40)

	list.traverse()
}
