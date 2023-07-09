// TODO Problem-10: Bubble Sort
// Bubble Sort is the simplest sorting algorithm that works 
// by repeatedly swapping the adjacent elements if they are in the wrong order.

let myArray = [8, 1, 2, 3, 4, 5, 6, 7]

function bubbleSort(array){
    let len = array.length;
    for(let i = 1; i < len; i++){
        console.log('i', i);
        console.log('Max j', len - i);

        let isSwapped = false;
        for(let j = 0; j < len - i; j++){
            if(array[j] > array[j+1]){
                [array[j], array[j+1]] = [array[j+1], array[j]];
                isSwapped = true;
            }
        }

        if(!isSwapped){
            break;
        }
    }
    return array;
}

let result = bubbleSort(myArray);
console.log('Result: ', result);