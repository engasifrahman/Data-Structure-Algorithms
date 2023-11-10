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
func (list *CircularLinkedList) deleteNode(data int) bool {
	if list.head == nil {
		fmt.Println("Circular linked list is empty.")
		return false
	}

	deleted := false
	current := list.head
	var prev *Node

	// Find the node to be deleted and its previous node
	for {
		if current.data == data {
			// Handle the case when the node to be deleted is the only node in the list
			if current == list.head && current.next == list.head {
				list.head = nil
			} else if current == list.head {
				// Handle the case when the node to be deleted is the head of the list
				prev = list.head
				for prev.next != list.head {
					prev = prev.next
				}
				prev.next = current.next
				list.head = current.next
			} else {
				// General case: update the next pointer of the previous node to skip the current node
				prev.next = current.next
			}

			// Set deleted = true when the node is found and deleted
			deleted = true
			break
		}

		if current.next == list.head {
			break
		}

		prev = current
		current = current.next
	}

	return deleted
}

// Traverse the circular linked list
func (list *CircularLinkedList) traverse() {
	if list.head == nil {
		fmt.Println("Circular linked list is empty.")
		return
	}

	current := list.head
	for {
		fmt.Printf("%d ", current.data)
		current = current.next
		if current == list.head {
			break
		}
	}
	fmt.Println()
}

func main() {
	list := CircularLinkedList{}

	list.insertAtBeginning(10)
	list.insertAfter(list.head, 25)
	list.insertAtBeginning(20)
	list.insertAfter(list.head, 15)
	list.insertAtBeginning(30)
	list.insertAtBeginning(40)
	list.insertAtEnd(70)

	list.traverse()

	list.deleteNode(40)
	list.deleteNode(25)
	list.deleteNode(70)

	list.traverse()
}
