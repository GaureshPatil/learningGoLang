package main

import (
	"fmt"
)

func main() {

	var colors = []string{"red", "greenn", "blue"}
	fmt.Println(colors)

	colors = append(colors, "yellow")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	number := make([]int, 5, 5)
	number[0] = 1
	number[1] = 7
	number[2] = 32
	number[3] = 12
	number[4] = 156
	fmt.Println(number)

	number = append(number, 235)
	fmt.Println(number)

	fmt.Println(cap(number))

}
