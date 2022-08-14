// HackerRank Problems
// Problem Name: pageCount  

// int n: the number of pages in the book
var n = 6;
// int p: the page number to turn to
var p = 5;

// Returns int: the minimum number of pages to turn

function pageCount (n, p) {
    if(p > n){
        return 0;
    }

    var quotient = parseInt(n / 2);
    // console.log('quotient', quotient);

    if(quotient >= p){
        for(let i = 0; i <= Math.round(quotient / 2); i++){
            if(i == 0 && i+1 == p){
                return i;
            }
            else if((i*2) == p || (i*2)+1 == p){
                return i;
            }
        }
    } else{
        for(let i = 0; i <= Math.round(quotient / 2); i++){
            if(n % 2 === 0){
                if(i == 0){
                    if(n === p){
                        return i;
                    }
                } else if((n - (i * 2) == p) || ((n+1) - (i * 2) == p)){
                    return i;
                }
            } else{
                if(i == 0){
                    if(n === p || n-1 === p){
                        return i;
                    }
                } else if((n - (i * 2) == p) || ((n-1) - (i * 2) == p)){
                    return i;
                }
            }
        }
    }
}

const result = pageCount (n, p);
console.log('Result: ', result);