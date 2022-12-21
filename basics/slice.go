package main

import (
	"fmt"
	"reflect"
)

func main() {
	/* Define an array containing programming languages */
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
	}

	/* Define slices */
	classics := languages[0:3]  // alternatively languages[:3]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[3:7]     // include 3 exclude 7
	new := languages[7:9]       // alternatively languages[7:]

	fmt.Printf("classic languagues: %v\n", classics) // classic languagues: [C Lisp C++]
	fmt.Printf("modern languages: %v\n", modern)     // modern languages: [Java Python JavaScript Ruby]
	fmt.Printf("new languages: %v\n", new)           // new languages: [Go Rust]

	/* MORE PLAYING WITH SLICE  */

	allLangs := languages[:]                     // copy of the array
	fmt.Println(reflect.TypeOf(allLangs).Kind()) // slice

	/* Create a slice containing web frameworks */
	// NOTICE: the square bracket is empty, if we pass parameter in it then it will be array
	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4]         // length 4 capacity 4
	frameworks = append(frameworks, "Meteor") // not possible with arrays
	/*
		The append function pushes new values to the end of a slice and returns a new slice with the same type as the original.
		In case the capacity of a slice is insufficient to store the new element, a new slice is created that can fit all the
		elements. In that case, the returned slice will refer to a different underlying array. Running the above code leads to
		the output below
	*/

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", jsFrameworks)
}
