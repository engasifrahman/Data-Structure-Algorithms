// TODO Problem-3: Count unique numbers
// !Input: [1, 1, 2, 3, 4, 4, 5, 6, 7, 8, 8]
// ?Solution: ?

//* My solutions
function countUnique(array){
    uniqueObject = {};

    for(number of array){
        if(uniqueObject[number]){
            continue;
        }
        console.log('number', number);

        uniqueObject[number] = 1;
        console.log('uniqueObject', uniqueObject);
    }

    console.log('Total Unique Numbers: ', Object.keys(uniqueObject).length);
    return Object.keys(uniqueObject).length;
}

// Tutorials solutions
function countUnique2(array){
    if(array.length){
        var i = 0;
        for(let j = 0; j < array.length; j++){
            if(array[i] !== array[j]){
                i++;
                array[i]=array[j];
            }
        }

        console.log('Total Unique Numbers: ', i+1);
        return i+1;
    } else {
        throw new Error('Array is empty!')
    }
}

countUnique2([1, 1, 2, 3, 4, 4, 5, 6, 7, 8, 8]);