// creating a stack in go using slices (not arrays). Use's go's implementation of resizing slices instead of manually resizing like in Java.
package main

import "fmt"

type ArrayStack[T any] struct {
	stack []T // create a slice called stack which is of any type
}

// Method to check if the stack is empty
func (s *ArrayStack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

// Method to push an item onto the stack
func (s *ArrayStack[T]) Push(item T) {
	s.stack = append(s.stack, item) // append handles resizing automatically
}

// Method to pop an item from the stack
func (s *ArrayStack[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue // return the zero value of the item if the stack is empty
	}

	item := s.stack[len(s.stack)-1]    // store the last element
	s.stack = s.stack[:len(s.stack)-1] // removes the last element
	return item
}

func main() {
	var words ArrayStack[string] // no need for a constructor
	words.Push("Food")
	words.Push("Water")
	fmt.Println(words.Pop())
}
