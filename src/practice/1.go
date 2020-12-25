package main

import (
	"fmt"
)

func main() {
	// creating map with key and value and both type is string eg. map[key] value and printing the result in next line
	countries := make(map[string] string)
		fmt.Println(countries)
		// assigning key and value to maps
		countries["IN"] = "India"
		countries["US"] = "United States"
		countries["BD"] = "Bangladesh"
		countries["AU"] = "Australia"

		fmt.Println(countries)

		// Iterating through maps using loop where k represents key and v represents the value and printing it in the loop
	for k, v := range countries {
		fmt.Println("Key:",k, "Value:", v)
	}
}
