package main

import "fmt"

/*
	the type `*T` is a pointer to a `T` value. The default value for pointers is nil.
*/

func main() {
	var address *int  // declare an int pointer
	number := 43      // int
	address = &number // address stores the memory address of number
	value := *address // dereferencing the value

	fmt.Printf("address: %v\n", address) // address: 0xc0000ae008
	fmt.Printf("value: %v\n", value)     // value: 42
}

/*
When working with pointers there are two important symbols to be aware of.
The address operator (&) provides the memory address of a value. It is used to bind a pointer to a value.
The asterisk operator (*) prefixing a type denotes a pointer type, whereas an asterisk prefixing a variable
is used to dereference the value the variable points to.
*/
