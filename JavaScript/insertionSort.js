// TODO Problem-112: Insertion Sort

// Insertion sort is a simple sorting algorithm that works similar to the way you sort playing cards in your hands. 
// The array is virtually split into a sorted and an unsorted part. 
// Values from the unsorted part are picked and placed at the correct position in the sorted part.

// Algorithm 
// To sort an array of size n in ascending order: 
// 1: Iterate from arr[1] to arr[n] over the array. 
// 2: Compare the current element (key) to its predecessor. 
// 3: If the key element is smaller than its predecessor, compare it to the elements before. 
// Move the greater elements one position up to make space for the swapped element.
let myArray = [8, 2, 5, 1, 9, 7];

function insertionSort(array){
    for(let i = 1; i < array.length; i++){
        let currentVal = array[i];
        let j = i - 1;
        
        while(j >= 0 && array[j] > currentVal){
            array[j+1] = array[j];
            j--;
        }
        array[j+1] = currentVal;
    }
    return array;
}

const result = insertionSort(myArray);
console.log('Result: ', result);