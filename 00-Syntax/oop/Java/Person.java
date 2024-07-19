package oop.Java; // must delcare the path this java file is in when doing oop. 

public class Person {
  private String name; // private field name
  private int age; // private field age
  static String species = "Human"; // static field species. This var is acesseble without creating an object and is the same no matter the object 

  public Person(String name, int age) { // public constructor api
    this.name = name;
    this.age = age;
  }

  // a function for an object is called a method. This is the greet method. 
  public void greet() { 
    System.out.println("Hello my name is " + name + ". I am " + age + " years old!");
  }

  public static String getSpecies(){ // static method that can be called without creating an object. Returns thestatic var species. 
    return species;
  }
}