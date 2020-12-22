package main

import (
	"fmt"
	"stringut"
)

func main() {

	n1, l1 := stringut.FullName("Gauresh", "Patil")
	fmt.Printf("Full Name: %v, number of chars: %v\n\n", n1, l1)

	n2, l2 := stringut.FullNameNakedReturn("Pritesh", "Patil")
	fmt.Printf("Full Name: %v, number of chars: %v\n\n", n2, l2)

}

// func fullName(f, l string) (string, int) {

// 	full := f + " " + l
// 	length := len(full)
// 	return full, length
// }

// func fullNameNakedReturn(f, l string) (full string, length int) {

// 	full = f + " " + l
// 	length = len(full)
// 	return
// }
