package main

import "fmt"

// go is a functional language not an OOP language. Uses structs to create new data types and calls functions/methods only. No classes or objects.

func main() {

	person1 := Person{ // this is how to create a Person struct
		name: "Bob",
		age:  24,
	}

	person1.Greet()
	person1.Greet2()
	fmt.Println()

	person1.IncreaseAge()
	person1.Greet2()
	fmt.Println()

	person1.setName()
	person1.Greet()
}
