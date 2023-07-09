// TODO Problem-2: String Anagram
// !String1: 'cinema'
// !String2: 'iceman'

// Small Changes

// *Solution
function isAnagram(string1, string2){
    var counter = {}
    if(string1.length !== string2.length){
        return false;
    }

    for(letter of string1){
        counter[letter] = (counter[letter] || 0)+1;
    }
    console.log(`counter`, counter);

    for(letter of string2){
        if(!counter[letter]){
            return false;
        }
        counter[letter] -= 1;
    }
    return true;
}

var result = isAnagram('cinema', 'iceman');
console.log(`result`, result);