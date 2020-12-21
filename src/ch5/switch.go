package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(6) + 1
	// fmt.Println("Day", dow)

	result := ""

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "It's Sunday"

	case 7:
		result = "It's Saturday"

	default:
		result = "Its a week day"
	}

	// fmt.Println("Day", dow, ",", result)
	fmt.Println(result)

	x := -42

	switch {
	case x < 0:
		result = "Less than zero"
		// fallthrough
	case x == 0:
		result = "equal to zero"

	default:
		result = "greater than zero"
	}
	fmt.Println(result)
}
