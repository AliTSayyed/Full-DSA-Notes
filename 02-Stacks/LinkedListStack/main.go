// create an empty stack using a linked list implementation
package main

import "fmt"

// Node represents an element in the linked list
type Node[T any] struct {
	item T        // stores any type
	next *Node[T] // has a pointer to the next node
}

// LinkedListStack represents a stack implemented using a linked list
type LinkedListStack[T any] struct {
	first *Node[T] // points to the first node in the stack
	size  int
}

// Method to check if the stack is empty
func (s *LinkedListStack[T]) IsEmpty() bool {
	return s.first == nil
}

// Method to push an item onto the stack
func (s *LinkedListStack[T]) Push(item T) {
	// copy the head
	oldFirst := s.first
	// update the current head
	s.first = &Node[T]{
		item: item,
		next: oldFirst,
	}
	// keep track of the size
	s.size++
}

// Method to pop an item off the stack
func (s *LinkedListStack[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	// remove the first item on the list
	item := s.first.item
	s.first = s.first.next
	s.size--
	return item
}

func main() {
	var words LinkedListStack[string] // Create a linkedListStack struct of type string
	words.Push("Food")
	words.Push("Water")
	fmt.Println(words.Pop())
}
