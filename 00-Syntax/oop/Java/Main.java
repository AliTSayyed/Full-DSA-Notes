package oop.Java; // must declare the path this java file is in when doing oop. 
// if importing a class from a different path (folder), use an import statement.
// since Person is in the same package (folder), no import statement is needed. 

public class Main {
  public static void main(String[] args) {
    Person person1 = new Person("Bob", 24); // create person object
    person1.greet(); // call object method 
    
    System.out.println(Person.getSpecies()); // can call the static person method without using a person object. 
  }
}
