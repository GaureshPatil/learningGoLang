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

	fmt.Println("the month is: ", t.Month())
	fmt.Println("the day is: ", t.Day())
	fmt.Println("the Weekday is: ", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v %v %v %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))

	shortFormat := "20/12/2020"
	fmt.Println("Tomorrow is ", tomorrow.Format(shortFormat))
}
