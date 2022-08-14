// TODO Problem-6: Check square elements of another array
// !Array1: [1, 2, 3, 4]
// !Array2: [1, 4, 9, 16]

//* Genaral Solutions
function checkSquare(array1, array2){
    for(let i =0; i < array1.length; i++){ 
        let isSquareChecked = false;

        for(let j =0; j < array2.length; j++){
            if(array1[i] * array1[i] == array2[j]){
                isSquareChecked = true;
            }

            if(j === (array2.length - 1) && !isSquareChecked){
                return false
            }
        }
    }

    return true;
}

// let result = checkSquare([1, 2, 3, 4], [3, 4, 9, 16]);

// console.log('Result: ', result);

//! Time complexity o(n^2) -> Quadratic

//* Optimized solutions

function checkSquare2(array1, array2){
    let map1 = {};
    let map2 = {};

    for(item1 of array1){
        map1[item1] = (map1[item1] || 0) + 1;
    }
    console.log('map1: ', map1);

    for(item2 of array2){
        map2[item2] = (map2[item2] || 0) + 1;
    }
    console.log('map2: ', map2);

    for(key in map1){
        // Obj key comapre in trem of square
        if(!map2[key * key]){
            return false;
        }

        // Value compare
        if(map1[key] !== map2[key * key]){
            return false;
        }
    }
    return true;
}

let result = checkSquare2([0, 2, 2, 4], [0, 4, 4, 16]);

console.log('Result2: ', result);
//* Time complexity o(n) -> liniar
