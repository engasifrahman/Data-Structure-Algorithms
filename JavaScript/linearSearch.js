// TODO Problem-9: Linear Search

// Linear search is a very simple search algorithm. 
// In this type of search, a sequential search is made over all items one by one. 
// Every item is checked and if a match is found then that particular item is returned, 
// otherwise the search continues till the end of the data collection.

let users = [
    { username: 'Asif', email: 'asif@gmail.com' }, 
    { username: 'Eva', email: 'eva@gmail.com' }, 
    { username: 'Yasha', email: 'yasha@gmail.com' }
];

function isUserExist(array, val){
    for(let item of array){
        if(item.username == val){
            return true;
        }
    }

    return false;
}

let result = isUserExist(users, 'Yasha2');
console.log('Result: ', result);