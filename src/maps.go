package main

import (
	"fmt"
)

func main() {

	states := make(map[string]string)
	fmt.Println(states)

	states["NGP"] = "Nagpur"
	states["MUM"] = "Mumbai"
	states["NMUM"] = "Navi-Mumbai"
	fmt.Println(states)

	Mumbai := states["MUM"]
	fmt.Println(Mumbai)

	delete(states, "NMUM")
	fmt.Println(states)

	states["NSK"] = "Nashik"

	for k, v := range states {
		fmt.Printf("%v %v\n", k, v)
	}
}
