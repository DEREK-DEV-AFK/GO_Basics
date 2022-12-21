package main

import "fmt"

func average(x []float64) (avg float64) {
	sum := 0.0
	switch len(x) {
	case 0:
		avg = 0
		//break NOTE: go does not need to use break
	default:
		for _, V := range x {
			sum += V
		}
		avg = sum / float64(len(x))

	}
	return
}

func main() {
	y := []float64{2.15, 3.14, 42.0, 29.5}

	fmt.Println(average(y))
}
