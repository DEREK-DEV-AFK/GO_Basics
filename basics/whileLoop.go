package main

import "fmt"

/*
In Go, there is no `while` keyword. Instead, we use the for keyword followed by a condition and a loop body.
The only exception is the missing semicolon at the end of the condition
*/

func main() {
	count := 0
	for count < 5 {
		fmt.Printf("hey count is less \n")
		count += 1
	}
}
