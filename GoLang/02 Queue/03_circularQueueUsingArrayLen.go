package main

import "fmt"

type Queue struct {
	items []int
	capacity int
}

func createQueue(capacity int) *Queue {
	return &Queue{
		items: make([]int, 0, capacity),
		capacity: capacity,
	}
}

// Add an element to the end of the queue
func (q *Queue) enqueue(item int) bool {
	if q.isFull() {
        fmt.Println("Queue is full, unable to push element:", item)
		return false
	}

	q.items = append(q.items, item)

	fmt.Printf("%d has enqueued to the Queue and updated Queue is %v \n", item, q.items)

	return true
}

// Remove an element from the front of the queue
func (q *Queue) dequeue() (int, bool) {
	item := 0

	if q.isEmpty() {
		fmt.Println("Unable to dequeue element as the queue is empty!")
		return item, false
	}

	item, q.items = q.items[0], q.items[1:]

	fmt.Printf("%d has dequeued from the Queue and updated Queue is %v \n\n", item, q.items)

	return item, true
}

// Get the value of the front of the queue without removing it
func (q *Queue) peek() (int, bool) {
	item := 0

	if q.isEmpty() {
		fmt.Println("There has no peek element as the queue is empty!")
		return item, false
	}

	item = q.items[0]

	fmt.Printf("%d is the peek item of the Queue and the Queue is %v \n", item, q.items)

	return item, true
}

// Check if the queue is empty
func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

// Check if the queue is full
func (q *Queue) isFull() bool {
	return len(q.items) == q.capacity
}

func main() {
	queue := createQueue(5)
	fmt.Println(queue.items)

    fmt.Println("Enqueuing elements to the Queue:")
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.enqueue(4)
	queue.enqueue(5)
	queue.enqueue(6) // Trying to enqueue element to a full Queue

    fmt.Println("\nDequeuing elements from the Queue:")
	queue.peek()
	queue.dequeue()

	queue.peek()
	queue.dequeue()

	queue.enqueue(6)

	queue.peek()
	queue.dequeue()

	queue.enqueue(7)

	queue.peek()
	queue.dequeue()

	queue.peek()
	queue.dequeue()

	queue.peek()
	queue.dequeue()

	queue.peek()
	queue.dequeue()

	queue.peek() // Trying to get peek element from an empty Queue
	queue.dequeue() // Trying to dequeue element from an empty Queue

	queue.enqueue(8)
}