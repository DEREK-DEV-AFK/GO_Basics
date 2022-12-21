package main

import (
	"fmt"
)

/* Declare single variable */
var a int

/* Declare a group of variables */
var (
	b bool
	c float32
	d string
)

func main() {
	fmt.Println("Hey there, this are types of variables")
	/* VARIABLES */
	a = 32           // assign single value
	b, c = true, 2.4 // assign multiple values
	d = "string"     // string must contain double quotes, IGNORE 😅

	/* CONSTANT */
	const e int = 12 // constant type, which can be declared in every data type

	/* SHORT DECLARING WAY */
	f := "hello"
	g := 34
	h, i := 3.4, true
	fmt.Println("int ->", a, "\n\tbool ->", b, "\n\t\t float ->", c, "\n\t\t\tstring ->", d, "\n\t\t\t\tshort-declare ->", f, g, h, i)

}
