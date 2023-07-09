// TODO Problem-1: Insertion Sort

// Bubble Sort is the simplest sorting algorithm 
// Insertion sort is a simple sorting algorithm that works similar to the way you sort playing cards in your hands. 
// The array is virtually split into a sorted and an unsorted part. 
// Values from the unsorted part are picked and placed at the correct position in the sorted part.

// Algorithm 
// To sort an array of size n in ascending order: 
// 1: Iterate from arr[1] to arr[n] over the array. 
// 2: Compare the current element (key) to its predecessor. 
// 3: If the key element is smaller than its predecessor, compare it to the elements before. 
// Move the greater elements one position up to make space for the swapped element.

package main
import ("fmt")

// *Time complexity of this InsertionSort is O(n^2) [Quadratic]
// *Space complexity of this InsertionSort is O(1)

func InsertionSort (array []int) []int {
	n := len(array)

	for i := 1; i < n; i++ {
		fmt.Println("Array:", array)

		item := array[i]

		j := i - 1
		for j >= 0 && array[j] > item {
			// Swap
			array[j + 1] = array[j]

			j--
		}

		array[j + 1] = item

	}

	return array;
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 9, 8, 7 }

	fmt.Println("Unsorted array:", arr)

	sortedArray := InsertionSort(arr)

	fmt.Println("Sorted array:", sortedArray)
}