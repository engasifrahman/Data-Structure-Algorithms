// TODO Problem: Find largest sum of consicutive digits [1, 2, 3, 4, 5, 6, 7, 8]

function findLargestSum(array, num){
    var max = 0;
    for(let i = 0; i < array.length - num + 1; i++){
        // console.log('i', i);

        var temp = 0;
        for(let j = 0; j < num; j++){
            // console.log('j', j);

            temp += array[i+j];
        }
        // console.log('temp', temp);

        if(temp > max){
            max = temp;
        }
    }

    console.log('max', max);
}

findLargestSum([1, 2, 3, 4, 5, 6, 7, 8], 4);
