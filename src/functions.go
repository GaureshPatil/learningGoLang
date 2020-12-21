package main

import (
	"fmt"
)

func main() {

	fmt.Println("")
	doSomething()

	sum := addValues(23, 54)
	fmt.Println(sum)

	sum = addAllValues(12, 54, 79, 80)
	fmt.Println("new sum:", sum)
}

func doSomething() {
	fmt.Println("doing")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
