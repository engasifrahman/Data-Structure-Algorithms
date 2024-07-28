package main

import "fmt"

func heapSort(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one
	for i := n - 1; i > 0; i-- {
		// Swap the root (maximum element) with the last element
		arr[0], arr[i] = arr[i], arr[0]

		// Heapify the reduced heap
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Find the largest element among the root, left child, and right child
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest is not the root, swap and heapify the affected sub-tree
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Original array:", arr)

	heapSort(arr)

	fmt.Println("Sorted array:", arr)
}
