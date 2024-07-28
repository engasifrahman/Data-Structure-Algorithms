package main

import "fmt"

type MaxHeap struct {
	array []int
	size  int
}

func createMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) insert(value int) {
	h.array = append(h.array, value)
	h.size++
	h.heapifyUp()
}

func (h *MaxHeap) heapifyUp() {
	index := h.size - 1
	for index > 0 && h.array[h.parent(index)] < h.array[index] {
		h.array[h.parent(index)], h.array[index] = h.array[index], h.array[h.parent(index)]
		index = h.parent(index)
	}
}

func (h *MaxHeap) extractMax() int {
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

func (h *MaxHeap) heapifyDown(index int) {
	maxIndex := index
	left := h.leftChild(index)
	right := h.rightChild(index)

	if left < h.size && h.array[left] > h.array[maxIndex] {
		maxIndex = left
	}

	if right < h.size && h.array[right] > h.array[maxIndex] {
		maxIndex = right
	}

	if index != maxIndex {
		h.array[maxIndex], h.array[index] = h.array[index], h.array[maxIndex]

		h.heapifyDown(maxIndex)
	}
}

func (h *MaxHeap) print() {
	for i := 0; i < h.size; i++ {
		fmt.Printf("%d ", h.array[i])
	}
	fmt.Println()
}

func main() {
	maxHeap := createMaxHeap()

	// maxHeap.insert(6)
	// maxHeap.insert(7)
	// maxHeap.insert(8)
	// maxHeap.insert(9)
	// maxHeap.insert(10)
	// maxHeap.insert(5)
	
	maxHeap.insert(10)
	maxHeap.insert(9)
	maxHeap.insert(8)
	maxHeap.insert(7)
	maxHeap.insert(6)
	maxHeap.insert(5)

	fmt.Print("Max-Heap array: ")
	// maxHeap.print()
	fmt.Printf("%v ", maxHeap.array)


	extractedMax := maxHeap.extractMax()
	fmt.Printf("Extracted max element: %d\n", extractedMax)

	fmt.Print("After extracting max element: ")

	maxHeap.print()
}
