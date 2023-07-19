package main
import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertAtBeggening(data int) {
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

func (list *LinkedList) InsertAtEnd(data int) {
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

func (list *LinkedList) InsertAfter(node *Node, data int) {

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

func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("List is empty");
		return
	}

	current := list.head

	for current != nil {
		if current.prev != nil {
			fmt.Println("Previous data:", current.prev.data)
		}

		fmt.Println("Current data:", current.data)
		fmt.Println("\n------------- ")


		current = current.next
	}
}

func main() {
	list := LinkedList{}

	list.InsertAtBeggening(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtBeggening(70)
	list.InsertAtEnd(40)
	list.InsertAtEnd(50)
	list.InsertAfter(list.head, 60)


	list.Display()
}