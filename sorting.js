// TODO Problem-7: Check sort
// !Array1: [2, 3, 1, 4]
let sortedArray = [];
let i = 0;
let j = 1;

function isSorted(array){
    for(let i=0; i< array.length; i++){
        if(array[i] > array[i+1]){
            return false;
        }
    }

    return true;
}

function checkSort(array){
    var is_sorted = isSorted(array);
    // console.log('isSorted(array)', is_sorted);

    if(is_sorted){
        sortedArray = array;
        return;
    }
    else if(array[i] < array[j]){
        i++;
        j++
        checkSort(array);
    } else{
        // [array[i], array[j]] = [array[j], array[i]];
        var b = array[j];
        array[j] = array[i];
        array[i] = b;

        i = 0;
        j = 1;
        checkSort(array);
    }
}

checkSort([2, 3, 1, 4]);
console.log('Sorted Array: ', sortedArray);