package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert at the beginning of the linked list
func (list *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{
		data: data, 
		next: list.head,
	}

	list.head = newNode
}

// Insert at the end of the linked list
func (list *LinkedList) InsertAtEnd(data int) {
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
func (list *LinkedList) InsertAfter(node *Node, data int) {
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
func (list *LinkedList) DeleteNode(data int) {
	if list.head == nil {
		return
	}

	// Case 1: If the node to be deleted is the head node
	if list.head.data == data {
		list.head = list.head.next
		return
	}

	// Case 2: If the node to be deleted is not the head node
	current := list.head
	prev := current
	for current != nil {
		if current.data == data {
			prev.next = current.next
			return
		}
		
		prev = current
		current = current.next
	}
}

// Display (traverse) the linked list
func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.InsertAtEnd(10)
	list.InsertAtBeginning(20)
	list.InsertAtBeginning(30)
	list.InsertAtEnd(40)
	list.InsertAfter(list.head.next, 50)

	fmt.Print("Linked list: ")
	list.Display()

	list.DeleteNode(40)

	fmt.Print("Linked list after deleting 30: ")
	list.Display()
}
