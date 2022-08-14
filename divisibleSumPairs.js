// HackerRank Problems

function divisibleSumPairs(n, k, ar) {
    // My Solution
    var total_possible_pairs = 0;
    var pair_sum = -1;
    var remainder = -1;    
    
    console.log('n: ', n);
    console.log('k: ', k);
    console.log('ar: ', ar);
    
    ar.sort(function(a, b) {
        return a - b;
    });
    console.log('Sorted ar: ', ar);
    
    for(let i = 0; i < n-1; i++){
       for(let j = i+1; j < n; j++){
            console.log('---------------------------');
            console.log('ar[i]: ', ar[i]);
            console.log('ar[j]: ', ar[j]);           
            if(i < j){               
                pair_sum = ar[i]+ar[j];
                console.log('pair_sum: ', pair_sum);
                remainder = pair_sum % k;
                console.log('remainder: ', remainder);
                               
                if(pair_sum >= k && remainder === 0){
                    total_possible_pairs++; 
                    console.log('This is a divisible pair');                  
                } else{
                    console.log('This is not a divisible pair');                               
                }
           }
       }
    }
    
    return total_possible_pairs;

    // Internet solution
    // var count = 0; 

    // for(let i = 0; i< n; i++){
    //     ar.slice(i+1, n).filter((item)=>{
    //         if((item + ar[i])%k === 0){
    //             count++
    //         }
    //     })
    // }

    // return count 
}

const n = 100;
const k = 22;
const ar = [43, 95, 51, 55, 40, 86, 65, 81, 51, 20, 47, 50, 65, 53, 23, 78, 75, 75, 47, 73, 25, 27, 14, 8, 26, 58, 95, 28, 3, 23, 48, 69, 26, 3, 73, 52, 34, 7, 40, 33, 56, 98, 71, 29, 70, 71, 28, 12, 18, 49, 19, 25, 2, 18, 15, 41, 51, 42, 46, 19, 98, 56, 54, 98, 72, 25, 16, 49, 34, 99, 48, 93, 64, 44, 50, 91, 44, 17, 63, 27, 3, 65, 75, 19, 68, 30, 43, 37, 72, 54, 82, 92, 37, 52, 72, 62, 3, 88, 82, 71]
var result = divisibleSumPairs(n, k, ar);
console.log('Result: ', result);