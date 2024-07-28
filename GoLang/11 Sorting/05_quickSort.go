package main

import "fmt"

func partition(arr []int, start, end int) int {
    // Choose the rightmost element as the pivot
    pivot := arr[end]

    // Initialize the index of the smaller element
    i := start - 1

    for j := start; j < end; j++ {
        // If the current element is less than or equal to the pivot
        if arr[j] <= pivot {
            i++
            // Swap arr[i] and arr[j]
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
	fmt.Println("i:", i)
	fmt.Println("end:", end)


    // Swap the pivot element with the element at index (i+1)
    arr[i+1], arr[end] = arr[end], arr[i+1]
	fmt.Println("partrition2:", arr)


    // Return the index of the pivot element
    return i + 1
}

func quickSort(arr []int, start, end int) {
    if start < end {
		// Partition the array into two parts and get the pivot index
		fmt.Println("---------------------------")
		fmt.Println("Start:", start)
		fmt.Println("End:", end)
		fmt.Println("Before partision:", arr)
		pivotIndex := partition(arr, start, end)
		fmt.Println("pivotIndex:", pivotIndex)

		fmt.Println("After partision:", arr)

		// Recursively sort the two sub-arrays
		quickSort(arr, start, pivotIndex-1)
		quickSort(arr, pivotIndex+1, end)
	}

}

func main() {
    // Test the Quick Sort algorithm
    arr := []int{12, 4, 5, 6, 7, 3, 1, 15}

    fmt.Println("Unsorted array:", arr)
    quickSort(arr, 0, len(arr)-1)

    fmt.Println("Sorted array:", arr)
}
