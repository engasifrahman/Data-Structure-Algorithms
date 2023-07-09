// HackerRank Problems
// Problem Name: sampleProblem 

// int x: Cat 's position
var x = 22;
// int y: Cat 's position
var y = 50;
// int z: Mouse 's position
var z = 9;

function sampleProblem(x, y, z) {
    var cat_a_distance = Math.abs(x-z);
    var cat_b_distance = Math.abs(y-z);

    console.log('cat_a_distance: ', cat_a_distance);
    console.log('cat_b_distance: ', cat_b_distance);

    if(cat_a_distance < cat_b_distance){
        return 'Cat A';
    } else if(cat_a_distance > cat_b_distance){
        return 'Cat B';
    } else{
        return 'Mouse C';
    }
}

const result = sampleProblem(x, y, z);
console.log('Result: ', result);