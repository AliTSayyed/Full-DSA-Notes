import { Person } from './person.js'
// need to tell Node.js to interpret your .js files as ES modules so create a package.json with type:module

class Main {
  static main() { // make the main class static so you do not need to create a main object to call the main class. 
    const person1 = new Person("Bob", 24);
    person1.greet();
  }
}

Main.main();

// second way of creating an object that does not require classes and imports 
const PersonObject = {
  name: "Bob",
  age: 30,
  weight: 222,
  greet: function() { // use an arrow function if you do not need to use instance vars, in this case we need to so use the normal method: function() call. 
    console.log(`Hello my name is ${this.name} and my age is ${this.age} and my weight is ${this.weight} lbs.`)
  }
}

PersonObject.greet();