package main

import (
	"fmt"
)

func main() {

	var x float64 = 42
	var result string

	if x < 0 {
		result = "less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)
}
