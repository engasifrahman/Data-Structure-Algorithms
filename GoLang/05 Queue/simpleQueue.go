/* ------------------------------ Simple Queue ------------------------------ */
// In a simple queue, insertion takes place at the tail and removal occurs at the head. It strictly follows the FIFO (First in First out) rule.

package main

import "fmt"

type Queue struct {
	data []int
	capacity, head, tail int
}

func newQueue(capacity int) Queue {
	return Queue{
		data: make([]int, capacity),
		capacity: capacity,
		head: 0,
		tail: 0,
	}
}

// Add an element to the end of the queue
func (q *Queue) enqueue(item int) bool {
	if q.isFull() {
        fmt.Println("Queue is full, unable to push element:", item)
		return false
	}

	q.data[q.tail] = item
	q.tail += 1

	fmt.Printf("%d has enqueued to the Queue and updated Queue is %v \n", item, q.data)

	return true
}

// Remove an element from the front of the queue
func (q *Queue) dequeue() (int, bool) {
	item := 0

	if q.isEmpty() {
        fmt.Println("Queue is empty, unable to dequeue element.")
		return item, false
	}

	item = q.data[q.head]
	q.data[q.head] = 0
	q.head += 1

	fmt.Printf("%d has dequeued from the Queue and updated Queue is %v \n\n", item, q.data)

	return item, true
}

// Get the value of the front of the queue without removing it
func (q *Queue) peek() (int, bool) {
	if q.isEmpty() {
		fmt.Println("Queue is empty, unable to peek element.")
		return 0, false
	}

	item := q.data[q.head]

	fmt.Printf("%d is the peek item of the Queue and the Queue is %v \n", item, q.data)

	return item, true
}

// Check if the queue is empty
func (q *Queue) isEmpty() bool {
	return q.head == q.tail
}

// Check if the queue is full
func (q *Queue) isFull() bool {
	if q.isEmpty() {
		q.tail = 0
		q.head = 0

		return false
	}

	return q.tail == q.capacity
}

func main() {
	queue := newQueue(5)

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

	queue.peek()
	queue.dequeue()

	queue.peek()
	queue.dequeue()

	queue.peek()
	queue.dequeue()

	queue.peek() // Trying to get peek element from an empty Queue
	queue.dequeue() // Trying to dequeue element from an empty Queue

	queue.enqueue(7)
}