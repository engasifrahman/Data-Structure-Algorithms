package main

import "fmt"

type MinHeap struct {
	array []int
	size  int
}

func createMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MinHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MinHeap) insert(value int) {
	h.array = append(h.array, value)
	h.size++
	h.heapifyUp()
}

func (h *MinHeap) heapifyUp() {
	index := h.size - 1
	for index > 0 && h.array[h.parent(index)] > h.array[index] {
		h.array[h.parent(index)], h.array[index] = h.array[index], h.array[h.parent(index)]
		index = h.parent(index)
	}
}

func (h *MinHeap) extractMax() int {
	if h.size == 0 {
		return -1 // Heap is empty
	}

	maxValue := h.array[0]

	h.array[0] = h.array[h.size-1]
	h.array = h.array[:h.size-1]

	h.size--
	h.heapifyDown(0)

	return maxValue
}

func (h *MinHeap) heapifyDown(index int) {
	minIndex := index
	left := h.leftChild(index)
	right := h.rightChild(index)

	if left < h.size && h.array[left] < h.array[minIndex] {
		minIndex = left
	}

	if right < h.size && h.array[right] < h.array[minIndex] {
		minIndex = right
	}

	if index != minIndex {
		h.array[index], h.array[minIndex] = h.array[minIndex], h.array[index]

		h.heapifyDown(minIndex)
	}
}

func (h *MinHeap) print() {
	for i := 0; i < h.size; i++ {
		fmt.Printf("%d ", h.array[i])
	}
	fmt.Println()
}

func main() {
	minHeap := createMinHeap()

	minHeap.insert(3)
	minHeap.insert(4)
	minHeap.insert(9)
	minHeap.insert(5)
	minHeap.insert(2)

	fmt.Print("Min-Heap array: ")
	minHeap.print()

	extractedMax := minHeap.extractMax()
	fmt.Printf("Extracted min element: %d\n", extractedMax)

	fmt.Print("After extracting min element: ")

	minHeap.print()
}
