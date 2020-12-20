package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("GO launched at %s\n", t)

	now := time.Now()
	fmt.Printf("the time now is: %s\n", now)
}
