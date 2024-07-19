package main

import (
	"fmt"
	"math"
)

func main() {
	// instantiate a variable

	var x int // first way to declare a var
	x = 1

	y := 1 // second way to decalre a var (NOT DYNAMICALLY TYPED)
	message := "Hello"
	boolean := true
	double := 0.5

	// print statement
	fmt.Println("Hello World!")
	fmt.Println("The value of x is", x)   // uses int value
	fmt.Printf("The value of x is %v", x) // formated print

	// conditional statements
	num := 10
	if num > 5 {
		fmt.Println("num is greater than 5")
	} else if num < 5 {
		fmt.Println("num is less than 5")
	} else { // num == 5
		fmt.Println("num is equal to 5")
	}

	num2 := 10
	num3 := 40
	if num2 == 10 && num3 < 50 { // and operator
		fmt.Println("true")
	} else if num2 < 20 || num3 == 20 { // or operator
		fmt.Println("true")
	}

	// loops
	//for declare; condition; update{}
	i := 5
	for i > 0 { // this is the while loop
		if i == 3 {
			i-- // must decrement before skipping printing 3 to avoid infinite loop
			continue
		}
		if i == 1 {
			break
		}
		fmt.Println(i)
		i--
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// math operations
	add := 5 + 3
	substract := 5 - 3
	mult := 5 * 3
	divide1 := 5 / 3     // integer division, truncates decimal
	divide2 := 5.0 / 3.0 // floating point division
	remainder := 5 % 3
	exponent := math.Pow(5, 3) // 5 ^ 3
	result := (5 - 3) / 2 * 5  // follows pemdas, does parenthesis first, then divides by 2, then multiplies by 5

	// calling funciton
	fmt.Println(greet("Bob"))
}

// function
func greet(name string) (greeting string) { // func (declares new function), greet (function name), name string (paramter name paramter type), greeting string (return name, return type)
	return "Hello " + name
} // can have multiple returns in one function
