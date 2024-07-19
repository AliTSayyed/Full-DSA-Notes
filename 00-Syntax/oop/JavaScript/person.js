export class Person {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }

  greet() { // greet method, do not need to write func in front 
    console.log(`Hello my name is ${this.name} and I am ${this.age} years old!`) // using `` notation to insert var in the string 
  }
}
