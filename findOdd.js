// TODO Problem-8: Find odd numbers
//! Array: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

function findOdd(array){
    let result = [];

    function helperFunction(inputArray){
        if(inputArray.length === 0){
            return;
        }
        if(inputArray[0] % 2 !== 0){
            result.push(inputArray[0]);
        }

        helperFunction(inputArray.slice(1));
    }

    helperFunction(array);
    return result;
}

const res = findOdd([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
console.log('Result: ', res);