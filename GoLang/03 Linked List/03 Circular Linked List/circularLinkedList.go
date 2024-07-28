package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

// Insert at the beginning of the circular linked list
func (list *CircularLinkedList) insertAtBeginning(data int) {
	newNode := &Node{
		data: data,
		next: list.head,
	}

	if list.head == nil {
		newNode.next = newNode // Make the node point to itself to form a circular list
		list.head = newNode
		return
	}

	current := list.head
	for current.next != list.head {
		current = current.next
	}
	current.next = newNode
	list.head = newNode
}

// Insert at the end of the circular linked list
func (list *CircularLinkedList) insertAtEnd(data int) {
	newNode := &Node{
		data: data,
		next: list.head,
	}

	if list.head == nil {
		newNode.next = newNode // Make the node point to itself to form a circular list
		list.head = newNode
		return
	}

	current := list.head
	for current.next != list.head {
		current = current.next
	}
	current.next = newNode
}

// Insert after a given node in the circular linked list
func (list *CircularLinkedList) insertAfter(node *Node, data int) {
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

// Delete a node from the circular linked list
func (list *CircularLinkedList) deleteNode(node *Node) bool {
	if node == nil || list.head == nil {
		return false
	}

	// Case-1; If the node to be deleted is the head node
	if list.head == node {
		if node.next == list.head { // Case-1.1: If the node to be deleted is the only node in the list
			list.head = nil
		} else { // Case-1.2: If the node to be deleted is connected with both head and last node
			current := list.head
			for current.next != list.head {
				current = current.next
			}

			current.next = node.next
			list.head = node.next
		}
		return true
	}

	// Case-2: if the node to be deleted is the internal node
	isDeleted := false
	current := list.head
	for current.next != list.head {
		if current.next == node {
			current.next = node.next
			isDeleted = true
			break
		}

		current = current.next
	}
	return isDeleted
}

// Delete a node by data from the circular linked list
func (list *CircularLinkedList) deleteNodeByData(data int) bool {
	if list.head == nil {
		return false
	}

	// Case-1; If the node to be deleted is the head node
	if list.head.data == data {
		if list.head.next == list.head { // Case-1.1: If the node to be deleted is the only node in the list
			list.head = nil
		} else { // Case-1.2: If the node to be deleted is connected with both head and last node
			current := list.head
			for current.next != list.head {
				current = current.next
			}

			current.next = current.next.next
			list.head = current.next.next
		}
		return true
	}

	// Case-2: if the node to be deleted is the internal node
	isDeleted := false
	current := list.head
	for current.next != list.head {
		if current.next.data == data {
			current.next = current.next.next
			isDeleted = true
			break
		}

		current = current.next
	}
	return isDeleted
}

// Traverse the circular linked list
func (list *CircularLinkedList) traverse() {
	if list.head == nil {
		fmt.Println("Circular linked list is empty.")
		return
	}

	fmt.Print("Traversal: ")


	current := list.head
	for {
		fmt.Print(" →", current.data)
		if current.next == list.head {
			break
		}
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := CircularLinkedList{}

	list.insertAtBeginning(40)
	list.insertAfter(list.head, 60)
	list.insertAtBeginning(30)
	list.insertAfter(list.head.next, 50)
	list.insertAtBeginning(20)
	list.insertAtBeginning(10)
	list.insertAtEnd(70)

	list.traverse()

	list.deleteNode(list.head)
	list.deleteNode(list.head.next.next.next)
	list.deleteNode(list.head.next.next.next.next.next.next)
	list.deleteNodeByData(25)
	list.deleteNodeByData(70)

	fmt.Println("\n------- After delete operation -------")
	list.traverse()
}
