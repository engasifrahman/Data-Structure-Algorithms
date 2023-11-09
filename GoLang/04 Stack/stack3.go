package main

import "fmt"

type Stack []int

// Add an element to the top of a stack
func (s *Stack) push(item int) {
	*s = append(*s, item)
	fmt.Printf("%d has pushed to the stack and updated stack is %v \n", item, *s)
}

// Remove an element from the top of a stack
func (s *Stack) pop() (int, bool) {
	item := -1
	if s.isEmpty() {
		fmt.Println("Stack is empty, unable to pop element.")
		return item, false
	}

	index := len(*s) - 1
	item, *s = (*s)[index], (*s)[:index]

	fmt.Printf("%d has popped from the stack and updated stack is %v \n\n", item, *s)

	return item, true
}

// Get the value of the top element without removing it
func (s *Stack) peek() (int, bool) {
	if s.isEmpty() {
		fmt.Println("Stack is empty, unable to peek element.")
		return -1, false
	}

	index := len(*s) - 1
	item := (*s)[index]

	fmt.Printf("%d is the peek value of the stack and the stack is %v \n", item, *s)

	return item, true
}

// Check if the stack is empty
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func main() {
	var stack Stack // Declare a stack using a slice
	// stack := Stack{} // Alternative declaration

	inputItem := 10;

	for i:= 0; i < inputItem; i++ {
		stack.push(i + 1)
	}

	fmt.Println("------------------------")

	for i:= 0; i < inputItem; i++ {
		stack.peek()
		stack.pop()
	}
	
	stack.peek()
	stack.pop()

	fmt.Println("------------------------")

	for i:= 0; i < inputItem; i++ {
		stack.push(i + 1)
	}
}
