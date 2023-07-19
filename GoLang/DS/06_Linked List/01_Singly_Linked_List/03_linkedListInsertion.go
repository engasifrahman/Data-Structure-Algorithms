package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (list *LinkedList) Insert(data int) {
    newNode := &Node{
		data: data, 
		next: nil,
	}

    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }

        current.next = newNode
    }
}

func (list *LinkedList) Display() {
    current := list.head

    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
}

func main() {
    list := LinkedList{}

    list.Insert(1)
    list.Insert(2)
    list.Insert(3)

    fmt.Println("Linked List:")
    list.Display()
}
