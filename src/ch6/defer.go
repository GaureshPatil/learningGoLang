package main 

import (
	"fmt"
)

func main() {

	defer fmt.Println("Close the File")
	fmt.Println("Open the File")

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("Undefered statement")

	x := 1000
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("value of x before defer:",x)

}

func myFunc() {

	defer fmt.Println("Defered in the Function")
	fmt.Println("Not Defered in the Function")
}
