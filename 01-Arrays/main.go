package main

import "fmt"

func main() {
	// An array is used to store multiple values in a single variable (must be of the same type).
	// Arrays are of fixed size. Can not increase the amount of space they have after declaring.
	// In GO, arrays are not refrences values. They are a value type. The array variable will directly contain all the array's elements.
	// The array variable will contain the memeory adress of the first value in the array. The other values in the array are in a contigous block for efficiency.

	// first way of creating an array and then inputing elements
	var numbers [5]int

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	// second way of creating an array and inputing elements
	cars := [3]string{"Truck", "SUV", "Cedan"}

	// this is how to loop through an array and acess each value
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// this is how to update a element at a certain index. Overwrites pervious element.
	cars[0] = "Electric Truck"

	// for each loop to iterate through each item in a list (another way). The first value is the index, can be ignored with an "_"
	for _, car := range cars {
		fmt.Println(car)
	}

	// this is how to get the size of an array. It is a property of the array.
	fmt.Println(len(numbers))

	// this is how to print out the entire array in one line. Unlike java, the array variable is the memeory adress of the first value in the list.
	// thats how it prints the array without needing a toString() method. In java, the array variable contains a value that points to the memeory adress of the first value in the list. Not the memeory adress of the first value itself.
	fmt.Println(cars)

	// You can create a copy of an array and modify it without affecting the original array. This is because values are not passed as refrences.
	numbers2 := numbers
	numbers2[4] = 10
	fmt.Println(numbers)
	fmt.Println(numbers2) // this is an entire different value with different memeory adresses. Not the same as java.

	var number3 *[5]int  // create a pointer for an array int of size 5
	number3 = &numbers   // assing the pointer to the original arrays memeory adress
	(*number3)[4] = 10   // modifying and derefrencing the pointer will affect the original array.
	fmt.Println(numbers) // the original array will now be modified.

	/* --------------------------------------------------------------------------------------------------------------------------------------------- */

	// a slice is what is mainly used in GO not arrays.
	// a slice is a dynamicaly sized flexible view into an array
	// slices are refrences to arrays

	// decalre and initilize array
	array := [6]int{1, 2, 3, 4, 5, 6}

	// this is how you declare a slice. Leave the brackets empty (no capacity declared)
	var slice []int

	// slices are built ontop of arrays
	slice = array[1:4]  // It contians values from index 1 to index 3 (does not include index 4). Last number is exclusive.
	slice2 := array[:]  // all elments
	slice3 := array[1:] // from index 1 - 5
	slice4 := array[:5] // from index 0 - 4

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	// you can still modify the original array using a pointer to a slice.
	var modArray *[]int // creates int slice pointer
	modArray = &slice   // sets pointer to memory adress of the slice
	(*modArray)[0] = 99 // derfrences pointer and modifies original array
	fmt.Println(array)  // printing original array will be modified.

	// this is how to create a slice without needing to explicitly create an array
	// slices created with the make will be filled with the zero value of the type
	// make(declare type, length, array capacity)
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) // use same len func to get size
	fmt.Println(cap(mySlice)) // can also get the capacity of the slice as well (if declared seperately)

	// usually written without mentioning an array capacity
	// default underlying array capacity will be the length of the slice (5 in this case)
	mySlice2 := make([]int, 5)
	fmt.Println(mySlice2)

	// third way to declare a slice with initilized values.
	// simply do not include a capacity in the [] and go will automatically consider it as a slice with an underlying array
	friends := []string{"Bob", "Sam", "Tim"}
	fmt.Println(friends)

	friends2 := friends
	friends2[0] = "Bobby"
	fmt.Println(friends)

	var modFriends *[]string
	modFriends = &friends
	fmt.Println(*modFriends) // the underlying array will be modified since friends2 is refrence of friends and changes the original slice and array. (like above)

	// this is how to add items to a slice
	friends = append(friends, "John")
	fmt.Println(friends)

	// Remove element by appending slices Need to create a new slice with a range excluding the value you do not want.
	// removed Sam from the list by updating the slice to include every element but Sam.
	// use the spread operator "..." to pass elments into a veriadic function (like append) to sort of "undo" the slices when processing in the append function
	friends = append(friends[:1], friends[2:]...)

	// can still iterate through a slice like an array
	for index, friend := range friends {
		fmt.Printf("%v: %s\n", index, friend)
	}

}
