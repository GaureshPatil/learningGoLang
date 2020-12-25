package main

import (
	"fmt"
	"time"
)

func main() {
	// To convert time to string
	t := time.Now()
	fmt.Printf("Time in String is : ",t.String())
	
	// any string to time 
	const layout = "2006-Jan-02"
	fmt.Printf("\n")
	tm, _ := time.Parse(layout, "2020-Dec-24")
	fmt.Println(tm)
}