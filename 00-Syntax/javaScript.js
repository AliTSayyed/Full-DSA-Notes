// instaitate a variable (dynamically typed)
const x = 1;
const message = 'hello';
const bool = true;
const double = 0.5;

// print statement
console.log('Hello World!');
console.log('The value of x is ', x); // keeps x as an int
console.log('The value of x is ' + x); // converts x to a string

// conditional statements
const num = 10;
if (num > 5) {
  console.log("num is greater than 5");
} else if (num < 5) {
  console.log("num is less than 5");
} else { // num === 5
  console.log("num is equal to 5");
}
const num2 = 10;
const num3 = 40;
if (num2 === 10 && num3 < 50) { // and operator, use triple === for equality comparison
  console.log("true");
} else if (num2 < 20 || num3 === 20) { // or operator
  console.log("true");
}
// 3 short cut operators (ternerary, gaurd, and default). Saves you writing out some code.      
// ternerary operator '?''
const result = true ? 'truthy' : 'falsy'; // if the boolean is true, the first variable is saved to const, if the boolean is false, the second value is saved to const
console.log(result);
// short cut for a if else statement saving a value to a variable based on a true or false condition. 
// gaurd operator && is similar to the AND operator but is used outside an if statement. If one part of the line is false, exits that line of code immediately. 
const text = false && console.log('message');// wont run because of false and guard operator. Short cut for if statement setting a variable to a certain value. 
// default operator sets the value of a variable to whichever value is true. If both are true the first value is saved
const currency = 'EUR' || 'USD';
console.log(currency) // will out put EUR but if EUR was replaced by undefinded, it would output USD. Sets a default statement.    

// loops 
let i = 5; // use key word let if variable's value will change/get updated
while (i > 0) {
  if (i === 3) {
    i--; // must decrement before skipping printing 3 to avoid infinite loop
    continue;
  }
  if (i === 1) {
    break;
  }
  console.log(i)
  i--;
}

for (let j = 0; j < 5; j++) {
  console.log(j);
}

// math operations
const add = 5 + 3;
const substract = 5 - 3;
const mult = 5 * 3;
const divide1 = 5 / 3; // automatic floating point divison
const divide2 = Math.floor(5 / 3); // integer division, rounds down
const remainder = 5 % 3;
const exponent1 = Math.pow(5, 3); // 5 ^ 3
const exponent2 = 5 ** 3;
const answer = (5 - 3) / 2 * 5; // follows pemdas, does parenthesis first, then divides by 2, then multiplies by 5

// functions

function greet1(name) { // stand alone function called greet 
  console.log('Hello ' + name);
}

const greet2 = (name) => { // arrow function version saved in a variable called greet
  console.log('Hello ' + name)
}

greet1("bobby")
greet2("bob"); 