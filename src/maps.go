package main

import (
	"fmt"
	"sort"
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

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
