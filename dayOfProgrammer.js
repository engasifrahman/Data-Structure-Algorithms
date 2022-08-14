// HackerRank Problems
// Problem Name: dayOfProgrammer 

const year = 2017;

function dayOfProgrammer(year) {
    var firstEightMonthDays = 0;
    if(year > 1918){
       if((year % 400 === 0) || (year % 4 === 0 && year % 100 !== 0)){
            firstEightMonthDays = 215 + 29;
        } else {
            firstEightMonthDays = 215 + 28;
        } 
    } else if(year < 1918){
        if((year % 4 === 0)){
            firstEightMonthDays = 215 + 29;
        } else {
            firstEightMonthDays = 215 + 28;
        } 
    } else{
        firstEightMonthDays = 215 + 28 - 13;
    }

    var dayOfSeptemper = 256 - firstEightMonthDays;

    return `${dayOfSeptemper}.09.${year}`;
}

const result = dayOfProgrammer(year);
console.log('Result: ', result);