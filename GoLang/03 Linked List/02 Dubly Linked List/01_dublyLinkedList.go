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
func (list *DoublyLinkedList) insertAtBeginning(data int) {
	newNode := &Node {
		prev: nil,
		data: data,
		next: list.head,
	}

	if list.head != nil {
		list.head.prev = newNode
	}

	list.head = newNode
}

// Insert at the end of the doubly linked list
func (list *DoublyLinkedList) insertAtEnd(data int) {
	newNode := &Node {
		prev: nil,
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

	newNode.prev = current
	current.next = newNode
}

// Insert after a given node in the doubly linked list
func (list *DoublyLinkedList) insertAfter(node *Node, data int) {
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

// Insert before a given node in the doubly linked list
func (list *DoublyLinkedList) insertBefore(node *Node, data int) {
	if node == nil {
		fmt.Println("Invalid node provided!")
	}

	newNode := &Node{
		data: data,
		prev: node.prev,
		next: node,
	}

	if node.prev != nil {
		node.prev.next = newNode
	}

	node.prev = newNode
}

// Delete a node from the doubly linked list by data value
func (l *DoublyLinkedList) deleteNode(node *Node) bool {
	if node == nil || l.head == nil {
		return false
	}

	// Case-1: If the node to be deleted is the head node
	if l.head == node {
		if node.next != nil {
			node.next.prev = nil
		}

		l.head = node.next
		return true
	}

	// Case-2: If the node to be deleted is not the head node
	isDeleted := false
	current := l.head
	for current.next != nil {
		if current.next == node {
			if node.next != nil {
				node.next.prev = current
			}

			current.next = node.next
			isDeleted = true

			break
		}

		current = current.next
	}
	return isDeleted
}

// Delete a node by data from the doubly linked list by data value
func (list *DoublyLinkedList) deleteNodeByData(data int) bool {
	if list.head == nil {
		return false
	}
	
	// Case-1: If the node to be deleted is the head node
	if list.head.data == data {
		if list.head.next != nil {
			list.head.next.prev = nil
		}

		list.head = list.head.next

		return true
	}

	// Case-2: If the node to be deleted is not the head node
	isDeleted := false
	current := list.head
	for current.next != nil {
		if current.next.data == data {
			if current.next.next != nil {
				current.next.next.prev = current
			}

			current.next = current.next.next
			isDeleted = true

			break
		}

		current = current.next
	}
	return isDeleted
}

// Traverse forward prints the elements of the doubly linked list in forward order
func (list *DoublyLinkedList) traverseForward() {
	fmt.Print("Forward traversal: ")

	current := list.head
	for current != nil {
		fmt.Print(" →", current.data)
		current = current.next
	}
	fmt.Println()
}

// Traverse backward prints the elements of the doubly linked list in backward order
func (list *DoublyLinkedList) traverseBackward() {
	fmt.Print("Backward traversal: ")

	// Find the last node in the list
	current := list.head
	for current != nil && current.next != nil{
		current = current.next
	}
	
	// Traverse backward and print the elements
	for current != nil {
		fmt.Print(" →", current.data)
		current = current.prev
	}
	fmt.Println()
}


func main() {
	list := DoublyLinkedList{}

	list.insertAtEnd(70)
	list.insertAtBeginning(60)
	list.insertAtBeginning(30)
	list.insertAtEnd(80)
	list.insertBefore(list.head.next, 40)
	list.insertAfter(list.head.next, 50)

	list.traverseForward()
	list.traverseBackward()

	list.deleteNode(list.head.next)
	list.deleteNodeByData(70)

	fmt.Println("\n------- After delete operation -------")
	list.traverseForward()
	list.traverseBackward()
}