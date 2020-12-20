package main

import (
	"fmt"
)

func main() {

	var colors = []string{"red", "greenn", "blue"}
	fmt.Println(colors)

	colors = append(colors, "yellow")
	fmt.Println(colors)
}
