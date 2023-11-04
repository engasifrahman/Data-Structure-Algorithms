package main
import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

type DoublyLinkedList struct {
	head *Node
}

// Insert at the beginning of the doubly linked list
func (list *DoublyLinkedList) InsertAtBeggening(data int) {
	newNode := &Node{
		prev: nil,
		data: data,
		next: list.head,
	}

	if list.head != nil {
		list.head.prev = newNode
	}

	list.head = newNode;
}

// Insert at the end of the doubly linked list
func (list *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{
		prev: nil,
		data: data,
		next: nil,
	}

	if list.head == nil {
		list.head = newNode;
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	newNode.prev = current
	current.next = newNode
}

// Insert after a given node in the doubly linked list
func (list *DoublyLinkedList) InsertAfter(node *Node, data int) {

	if node == nil {
		fmt.Println("Given node is empty");
		return
	}

	newNode := &Node{
		prev: node,
		data: data,
		next: node.next,
	}

	if node.next != nil {
		node.next.prev = newNode
	} 

	node.next = newNode
}

// Delete a node from the doubly linked list by data value
func (list *DoublyLinkedList) DeleteNode(data int) {
	if list.head == nil {
		return
	}

	// Case 1: If the node to be deleted is the head node
	if list.head.data == data {
		if list.head.next != nil {
			list.head.next.prev = nil
		} 

		list.head = list.head.next
	}

	// Case 2: If the node to be deleted is not the head node
	current := list.head
	prev := current
	for current != nil {
		if current.data == data { 
			if current.next != nil {
				current.next.prev = prev
			}
			prev.next = current.next

		}

		prev = current
		current = current.next
	}
}

// DisplayForward prints the elements of the doubly linked list in forward order
func (list *DoublyLinkedList) DisplayForward() {
	fmt.Println("Doubly linked list (forward):")
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

// DisplayBackward prints the elements of the doubly linked list in backward order
func (list *DoublyLinkedList) DisplayBackward() {
	fmt.Println("Doubly linked list (backward):")
	// Find the last node in the list
	current := list.head
	for current != nil && current.next != nil {
		current = current.next
	}

	// Traverse backward and print the elements
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}
	fmt.Println()
}

func main() {
	list := DoublyLinkedList{}

	list.InsertAtBeggening(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtBeggening(70)
	list.InsertAtEnd(40)
	list.InsertAtEnd(50)
	list.InsertAfter(list.head, 60)

	list.DeleteNode(40)

	list.DisplayForward()
	list.DisplayBackward()
}