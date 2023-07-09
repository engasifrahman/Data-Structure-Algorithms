// HackerRank Problems
// Problem Name: migratoryBirds

const arr = [1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4];

function migratoryBirds(arr) {
    let sightedBirds = {};

    for(item of arr){
        // console.log('item', item);

        if(!sightedBirds[item]){
            sightedBirds[item] = 1;
        } else{
            sightedBirds[item] += 1;
        }
    }
    console.log('sightedBirds', sightedBirds);

    var max_count = Math.max(...Object.values(sightedBirds));
    console.log('max_count', max_count);

    var max_ids = Object.entries(sightedBirds).filter(function(innerArray, index){
        console.log('innerArray', innerArray);

        return innerArray[1] == max_count;
    });
    
    console.log('max_ids', max_ids);

    var smallest_id = Math.min(...Object.keys(Object.fromEntries(max_ids)))
    console.log('smallest_id', smallest_id);

    return smallest_id;
}

const result = migratoryBirds(arr);
console.log('Result: ', result);