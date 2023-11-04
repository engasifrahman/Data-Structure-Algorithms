// TODO Problem-2: Bubble Sort

// Bubble Sort is the simplest sorting algorithm 
// Bubble Sort works by repeatedly swapping the adjacent elements if they are in the wrong order. 
// This algorithm is not suitable for large data sets as its average and worst-case time complexity is quite high.

package main
import ("fmt")

// *Time complexity of this BubbleSort is O(n^2) [Quadratic]
// *Space complexity of this BubbleSort is O(1)

func BubbleSort (array []int) []int {
	n := len(array)

	for i := 1; i < n; i++ {
		fmt.Println("Array:", array)

        isSwapped := false;

		for j := 0 ; j < n - i; j++ {
			if array[j] > array[j + 1] {
				// Swap
				array[j], array[j + 1] = array[j + 1], array[j]
				isSwapped = true;
			}
		}

		if !isSwapped {
            break;
        }
	}

	return array;
}

func main() {
	arr := []int{ 9, 8, 7, 6, 5, 4, 3, 2, 1 }

	fmt.Println("Unsorted array:", arr)

	sortedArray := BubbleSort(arr)

	fmt.Println("Sorted array:", sortedArray)
}