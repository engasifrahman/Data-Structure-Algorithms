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
func (list *LinkedList) deleteNode(node *Node) bool {
	if node == nil || list.head == nil {
		return false
	}

	// Case-1: If the node to be deleted is the head node
	if list.head == node {
		list.head = node.next
		return true
	}

	// Case-2: If the node to be deleted is an internal node
	isDeleted := false
	current := list.head
	for current.next != nil {
		if current.next == node {
			current.next = node.next
			isDeleted = true
			break
		}

		current = current.next
	}
	return isDeleted
}

// Delete a node by data from the linked list
func (list *LinkedList) deleteNodeByData(data int) bool {
	if list.head == nil {
		return false
	}

	// Case-1: If the node to be deleted is the head node
	if list.head.data == data {
		list.head = list.head.next
		return true
	}

	// Case-2: If the node to be deleted is an internal node
	isDeleted := false
	current := list.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			isDeleted = true
			break
		}

		current = current.next
	}
	return isDeleted
}

// Traverse the linked list
func (list *LinkedList) traverse() {
	fmt.Print("Traversal: ")

	current := list.head
	for current != nil {
		fmt.Print(" →", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.insertAtEnd(60)
	list.insertAtBeginning(40)
	list.insertAtBeginning(30)
	list.insertAfter(list.head.next, 50)
	list.insertAtBeginning(20)

	list.traverse()

	list.deleteNode(list.head.next)
	list.deleteNodeByData(40)
	list.deleteNode(list.head.next.next)

	fmt.Println("\n------- After delete operation -------")
	list.traverse()
}
