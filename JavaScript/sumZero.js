// TODO Problem-1: Cheking sum zero
// !INPUT: [-5, -4, -3, -2, 0 , 2, 4, 6, 8]
// ?Output: [?, ?]

// *General Solution
function getGeneralSumPairZero(array){
    for(number of array){
        // console.log(`number`, number)

        for(value of array){
            // console.log(`value`, value);

            if(number+value === 0){
                // console.log(`number+value`, number+value);
                return [number, value];
            }
        }
    }

    return 'Not Found';
}
// !Time complexity of this getGeneralSumPairZero is o(n^2) [quadratic]

// *Optimised Solution
function getOptimisedSumPairZero(array){
    var result = [];
    var leftIndex = 0
    var rightIndex = array.length - 1;

    while(leftIndex < rightIndex){
        var sumPair = array[leftIndex]+array[rightIndex];
        // console.log(`array[leftIndex]`, array[leftIndex]);
        // console.log(`array[rightIndex]`, array[rightIndex]);
        // console.log(`sumPair`, sumPair);

        if(sumPair === 0){
            return [array[leftIndex], array[rightIndex]]
        }else if(sumPair < 0){
            leftIndex++;
        }else{
            rightIndex--;
        }
    }

    return 'Not Found';
}
// *Time complexity of this getOptimisedSumPairZero is o(n) [linear]

// var result = getGeneralSumPairZero([-5, -4, -3, -2, 0 , 2, 4, 6, 8]);
var result = getOptimisedSumPairZero([-5, -4, -3, -2, 0 , 2, 4, 6, 8]);
console.log(`result`, result);