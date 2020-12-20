package main

import (
	"fmt"
)

type Cat struct {
	Breed  string
	Weight int
}

func main() {

	doodle := Cat{"doodle", 20}
	fmt.Println(doodle)
	fmt.Printf("%+v", doodle)
}
