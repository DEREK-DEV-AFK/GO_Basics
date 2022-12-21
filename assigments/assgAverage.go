/*
We define an input slice called y in the body of the main function and we make a call to average,
passing in x as an argument. We wrap the call inside fmt.Println() to write the result to standard output.
The interesting part is the implementation of the average function. Notice the return parameter avg is defined
immediately at end of the function declaration. In the function body, we initialize a variable named total that
will compute a running sum of the slice elements. From there, we check the size of the input slice. If the slice
is empty, we return 0, otherwise, we loop through each element in the slice and add it to the total. Notice how we
use the underscore (_) for the unused variable. We convert the length of the slice to a float using float64(len(x)).
Finally, we compute the average and return the result to the caller.
*/

package main

import "fmt"

func average(x []float64) (avg float64) { // one in comment is also declaring the return variable and type
	sum := 0.0

	if len(x) == 0 {
		avg = 0
	} else {
		for _, value := range x { // we use `_` for unsed variables
			sum += value
		}

		avg = sum / float64(len(x))
	}
	return
}

func main() {
	y := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average(y)) // 19.197499999999998
}
