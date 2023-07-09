// Divide & conquerer technique
// TODO Problem5: Find index of given no in a sorted array [Given no is 7]
// !Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]
// ?Solutions

// Small Changes

function binarySearch(array, num){
    var minIndex = 0;
    var maxIndex = array.length - 1;

    while(minIndex <= maxIndex){
        console.log('-----------------------------------------------------');
        console.log('minIndex', minIndex);
        console.log('maxIndex', maxIndex);
    
        var midIndex = Math.floor((minIndex + maxIndex) / 2);
        console.log('midIndex', midIndex);

        if(array[midIndex] < num){
            minIndex = midIndex + 1;
            console.log('minIndex', minIndex);
        } else if(array[midIndex] > num){
            maxIndex = midIndex - 1;
            console.log('maxIndex', maxIndex);
        } else{
            console.log('Result is:', midIndex);
            return midIndex;
        }
    }
}

binarySearch([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], 7);
