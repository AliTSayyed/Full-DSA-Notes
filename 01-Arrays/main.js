// arrays in JavaScript are dynamic and can allow store different data types
// there is no fixed size array in java script
// array variables hold refrences adresses. Those refrence adress then point to the memory adress of the first elemetn in the list. Subsequent elemetents are stored in contgious blocks.

/*
Here is an example of a refrence:
array1 ---> Ref1234
array2  ---> Ref1234
Ref1234 ---> [1, 2, 3, 4, 5]
*/

let numbers = [1, 2, 3, 4, 5, 6];
// can access elements using their index 
console.log(numbers[0]);

// can modify indexes
numbers[0] = 10;

// arrays can store different types and objects. 
let mixedTypes = [1,2,"Dog", true];
console.log(mixedTypes);

let fruits = ['apple', 'orange', 'banana', 'coconut'];
console.log(fruits);

// this is how to add elements to the end of the list
fruits.push('coconut');

// how to remove elements from the end of the list
fruits.pop();

// how to add elements to the begining of the list 
fruits.unshift('mango');

// how to remove elements from the begining of the list
fruits.shift();

let numOfFruits = fruits.length // how to get length of an array

// how to search for the index of an element
let index = fruits.indexOf('apple');
console.log(index); // returns -1 if element is not found 

// use a for loop to iterate through the array
for (let i = 0; i < fruits.length; i++){
  console.log(fruits[i]); 
}

// for each loop
for (let fruit of fruits){
  console.log(fruit);
}


// another forEach loop (better)
fruits.forEach((fruit) => {
  console.log(fruit);
});

// how to sort the array
fruits.sort(); 
console.log(fruits);

// arrays are refrences as well. If you make a copy of an array you are making another variable but assigning it to the same memeory adress
// so chaning a copy of the array will affect the original array. 
let array2 = numbers;
array2.pop();
console.log(numbers);

// the splice method is used to add remove or replace elemetns in an array and returns the modified array
let arr1 = [1,2,3,4,5,6];
let removedArr1 = arr1.splice(2,2); // start at index 2 and delete 2 elements (the second element and one more). Store the deleted elments in modArr1.
console.log(arr1);
console.log(removedArr1);

arr1.splice(2,0,"a","b"); // start at index 2, remove 0 elements, add a and b
console.log(arr1);

let replacedArr1 = arr1.splice(2,2,3,4); // start at index 2, remove 2 elements (index 2 and the next element), and add numbers 3 and 4 starting from index 2. 
console.log(arr1);
console.log(replacedArr1); // holds the deleted items. 

// do not need to save the splice to a var can just call it on an array as well
arr1.splice(0,2);
console.log(arr1);