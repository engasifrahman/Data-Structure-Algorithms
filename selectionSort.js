// TODO Problem-11: Selection Sort

// The selection sort algorithm sorts an array by repeatedly finding the 
// minimum element (considering ascending order) from unsorted part and putting 
// it at the beginning. The algorithm maintains two subarrays in a given array.
// 1) The subarray which is already sorted. 
// 2) Remaining subarray which is unsorted.
// In every iteration of selection sort, 
// the minimum element (considering ascending order) from the unsorted subarray is picked and moved to the sorted subarray. 

function selectionSort(array){
    const arr = Array.from(array); // avoid side effects
    const len = arr.length;
    
    for(let i = 0; i < len - 1; i++) {
      let minPos = i;
      for(let j = i + 1; j < len; j++) {
        if(arr[j] < arr[minPos]) {
          minPos = j;
        }
      }
      if(i !== minPos) {
        [arr[i], arr[minPos]] = [arr[minPos], arr[i]];
      }
    }

    return arr;
  };
  
  console.log(selectionSort([4, 9, 2, 1, 5]));