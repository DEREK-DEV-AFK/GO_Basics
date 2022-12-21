package main

import "fmt"

/*
A struct is simply a collection of fields.
*/

// Define a stack type using the struct
type stack struct {
	index int
	data  [5]int
}

/* Define push and pop methods */
func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

/* Notice the stack pointer s passed as an argument */
func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}

func main() {
	/* Create a pointer to the new stack and push 2 values */
	s := new(stack)

	s.push(3)
	s.push(5)
	s.push(6)

	fmt.Println(s) // / stack: {3 [3 5 6 0 0]}
}
