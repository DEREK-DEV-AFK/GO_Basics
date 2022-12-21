package main

import "fmt"

func main() {
	/* Short way */
	a, b, c, d := 12, true, 3.45, "string"

	/* `%T` is called as verb for type to get the data type of variable */
	fmt.Printf("Data types are \n %T %T %T %T \n", a, b, c, d)

	/* `%d` is a verb for decimals */
	fmt.Printf("int : %d", a)

	/*
		`%s` - for string
		`%f`- for float
		`%t` - for boolean
		`%v` - for any natural value for a type
	*/
}
