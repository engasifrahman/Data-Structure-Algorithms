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
func (l *DoublyLinkedList) insertAtBeginning(data int) {
	newNode := &Node {
		prev: nil,
		data: data,
		next: l.head,
	}

	if l.head != nil {
		l.head.prev = newNode
	}

	l.head = newNode
}

// Insert at the end of the doubly linked list
func (l *DoublyLinkedList) insertAtEnd(data int) {
	newNode := &Node {
		prev: nil,
		data: data,
		next: nil,
	}

	if l.head == nil {
		l.head = newNode
		return	
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	newNode.prev = current
	current.next = newNode
}

// Insert after a given node in the doubly linked list
func (l *DoublyLinkedList) insertAfter(node *Node, data int) {
	if node == nil {
		return
	}

	newNode := &Node {
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
func (l *DoublyLinkedList) deleteNode(data int) bool {
	if l.head == nil {
		return false
	}
	
	// Case 1: If the node to be deleted is the head node
	if l.head.data == data {
		if l.head.next != nil {
			l.head.next.prev = nil
		}

		l.head = l.head.next

		return true
	}

	// Case 2: If the node to be deleted is not the head node
	removed := false
	current := l.head

	for current.next != nil {
		if current.next.data == data {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next

			removed = true
			break
		}

		current = current.next
	}

	return removed
}

// Traverse forward prints the elements of the doubly linked list in forward order
func (l *DoublyLinkedList) traverseForward() {
	fmt.Print("\nForward traversal: ")

	current := l.head
	for current != nil {
		fmt.Printf("%d →", current.data)
		current = current.next
	}
}

// Traverse backward prints the elements of the doubly linked list in backward order
func (l *DoublyLinkedList) traverseBackward() {
	fmt.Print("\nBackward traversal: ")

	// Find the last node in the list
	current := l.head
	for current != nil && current.next != nil{
		current = current.next
	}
	
	// Traverse backward and print the elements
	for current != nil {
		fmt.Printf("%d →", current.data)
		current = current.prev
	}
}


func main() {
	list := DoublyLinkedList{}

	list.insertAtBeginning(10)
	list.insertAtEnd(20)
	list.insertAtEnd(30)
	list.insertAtBeginning(70)
	list.insertAtEnd(40)
	list.insertAtEnd(50)
	list.insertAfter(list.head, 60)

	list.deleteNode(40)
	list.deleteNode(80)

	list.traverseForward()
	list.traverseBackward()
}