// TODO Problem-2: Binary Search

// Binary search is a common search algorithm
// Binary search is a search algorithm that finds the position of a target value within a sorted array. 
// Binary search compares the target value to the middle element of the array.

package main
import ("fmt")

// *Time complexity of this linearSearch is O(log n) [logarithmic]
// *Space complexity of this linearSearch is O(1)

func binarySearch (arr []int, search int) int {
	left := 0
	right := len(arr) - 1

	result := -1

	for left <= right {
		mid := int((left + right) / 2)

		if arr[mid] == search {
			result = mid
			break
		}

		if arr[mid] < search {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func main() {
	arr := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 }
    search := 7;

	searchindex := binarySearch(arr, search)

	fmt.Println(arr)
	fmt.Printf("In the above array the index of %v is %v \n", search, searchindex)
}