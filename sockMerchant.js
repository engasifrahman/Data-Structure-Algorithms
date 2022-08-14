// HackerRank Problems
// Problem Name: sockMerchant

const arr = [1, 2, 1, 2, 1, 3, 2];

function sockMerchant(arr) {
    let colorStocks = {};
    let totalPairs = 0;

    for(item of arr){
        // console.log('item', item);

        if(!colorStocks[item]){
            colorStocks[item] = 1;
        } else{
            colorStocks[item] += 1;
        }
    }
    console.log('colorStocks', colorStocks);

    for(item of Object.entries(colorStocks)){
        // console.log('item', item);

        var quotient = item[1] / 2;

        console.log('quotient', parseInt(quotient));

        totalPairs += parseInt(quotient);
    }

    return totalPairs
}

const result = sockMerchant(arr);
console.log('Result: ', result);