package main

// as long as files are in the same package (main path) no import statmeent is needed
// to import you use go get ..., go mod vendor, go mod tidy, go mod init, etc.

import "fmt"

type Person struct { // new data type called 'Person' that has a name (string type) and age (int type) data fields.
	name string
	age  int
}

func (p Person) Greet() { // function for a struct is called a method. Declare the struct type after 'func' to make a method
	fmt.Println("Hello my name is", p.name)
}

func (p Person) Greet2() { // uppercase functions/methods are acessible to files outside this package.
	fmt.Printf("Hello my name is %v and I am %v years old!", p.name, p.age)
}

func (p *Person) IncreaseAge() { // use pointers to modify original value
	p.age++
}

func (p *Person) setName() { // lower case functions are private to the package in which they are defined
	p.name = "Default name"
}
