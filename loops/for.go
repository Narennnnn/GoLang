package main

import (
	"fmt"
)

func main() {
	initialValue := 0
	for i := 0; i < 5; i++ {
		fmt.Println(initialValue)
		initialValue += i
	}
	fmt.Println(initialValue)
	fmt.Printf("The type: %T \n", initialValue)

	/*

		For is Go's "while"
		At that point you can drop the semicolons: C's while is spelled for in Go.
	*/
	for initialValue < 100 {
		initialValue += initialValue
	}
	fmt.Println(initialValue)
}
