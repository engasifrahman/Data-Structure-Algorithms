// TODO Problem-1: Linear Search

// Linear search is a very simple search algorithm. 
// In this type of search, a sequential search is made in over all items one by one. 
// Every item is checked and if a match is found then that particular item is returned, 
// otherwise the search continues till the end of the data collection.

package main
import ("fmt")

// *Time complexity of this linearSearch is O(n) [linear]
// *Space complexity of this linearSearch is O(1)

func linearSearch (arr []int, search int) int {
    for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			return i;
		}
	}
        
    return -1;
}

func main() {
	arr := []int{ 2, 3, 4, 10, 40, 6, 99, 33 }
    search := 99;

	searchindex := linearSearch(arr, search)

	fmt.Println(arr)
	fmt.Printf("In the above array the index of %v is %v \n", search, searchindex)
}

