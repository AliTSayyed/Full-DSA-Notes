// create an empty stack using a linked list implementation
package main

// Node represents an element in the linked list
type Node[T any] struct {
	item T        // stores any type
	next *Node[T] // has a pointer to the next node
}

// LinkedListStack represents a stack implemented using a linked list
type LinkedListStack[T any] struct {
	first *Node[T]
	size  int
}

// Method to check if the stack is empty
func (s *LinkedListStack[T]) IsEmpty() bool {
	return s.first == nil
}

func main() {

}
