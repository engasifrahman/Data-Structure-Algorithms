// TODO Problem-1: Selection Sort

// Selection sort is a simple and efficient sorting algorithm 
// Selection sort works by repeatedly selecting the smallest (or largest) element 
// from the unsorted portion of the list and moving it to the sorted portion of the list. 

// Procedures
// It sorts from left to right
// First loop will run n-1 times, and will start from 0th index to (n-1)th index
// Second loop will run maximum of [(n-1), (n-2), (n-3), (n-4), ..., 1] times, and will start from 1st index to nth index

package main
import "fmt"

// *Time complexity of the Selection Sort is O(n^2) [Quadratic]
// *Space complexity of the Selection Sort is O(1)

func selectionSort (array []int) []int {
	n := len(array)

	for i := 0; i < n - 1; i++ {
		min_idx := i
		for j := i + 1 ; j < n; j++ {
			if array[min_idx] > array[j] {
				min_idx = j
			}
		}

		if i != min_idx {
			// Swap
			array[i], array[min_idx] = array[min_idx], array[i]
		}
	}

	return array;
}

func main() {
	arr := []int{ 4, 9, 2, 1, 5 }

	fmt.Println("Unsorted array:", arr)

	sortedArray := selectionSort(arr)

	fmt.Println("Sorted array:", sortedArray)
}