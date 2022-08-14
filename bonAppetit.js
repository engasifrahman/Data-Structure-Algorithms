// HackerRank Problems
// Problem Name: bonAppetit 

const bill = [3, 10, 2, 9];
const k = 1;
const b = 12;

function bonAppetit(bill, k, b) {
    var total_bill = bill.reduce((partialSum, a) => partialSum + a, 0);

    if(bill[k]){
        total_bill -= bill[k];
    }

    var annas_part = total_bill / 2;

    var result = annas_part == b ? 'Bon Appetit' : b - annas_part;
    console.log(result);
}

bonAppetit(bill, k, b);