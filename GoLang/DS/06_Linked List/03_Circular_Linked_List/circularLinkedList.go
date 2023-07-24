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
func (list *CircularLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		newNode.next = newNode // Make the node point to itself to form a circular list
		list.head = newNode
	} else {
		current := list.head
		for current.next != list.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = list.head
		list.head = newNode
	}
}

// Insert at the end of the circular linked list
func (list *CircularLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		newNode.next = newNode // Make the node point to itself to form a circular list
		list.head = newNode
	} else {
		current := list.head
		for current.next != list.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = list.head
	}
}

// Insert after a given node in the circular linked list
func (list *CircularLinkedList) InsertAfter(prevNode *Node, data int) {
	if prevNode == nil {
		fmt.Println("Invalid node provided")
		return
	}

	newNode := &Node{data: data, next: nil}
	newNode.next = prevNode.next
	prevNode.next = newNode

	if prevNode == list.head { // Update head if inserting after the head node
		list.head = newNode
	}
}

// Delete a node from the circular linked list
func (list *CircularLinkedList) DeleteNode(data int) {
	if list.head == nil {
		fmt.Println("Circular linked list is empty.")
		return
	}

	current := list.head
	var prev *Node

	for {
		if current.data == data {
			if current == list.head { // If the node to be deleted is the head
				list.head = current.next
				if current.next == current { // If there's only one node in the list
					list.head = nil
				} else {
					prev.next = current.next
				}
			} else {
				prev.next = current.next
			}
			break
		}

		prev = current
		current = current.next

		if current == list.head { // Node not found after traversing the entire list
			fmt.Println("Node not found.")
			break
		}
	}
}

// Display (traverse) the circular linked list
func (list *CircularLinkedList) Display() {
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

	list.InsertAtBeginning(10)
	list.InsertAtBeginning(20)
	list.InsertAfter(list.head, 15)
	list.InsertAtBeginning(30)
	list.InsertAtBeginning(40)
	list.InsertAfter(list.head, 25)
	list.InsertAtEnd(70)


	fmt.Print("Circular linked list: ")
	list.Display()

	list.DeleteNode(30)

	fmt.Print("Circular linked list after deleting node with data 30: ")
	list.Display()
}
