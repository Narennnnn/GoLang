/*
package main

import "fmt"

	func main() {
		defer fmt.Println("world")

//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

		fmt.Println("hello")
	}
*/
package main

import "fmt"

func main() {
	fmt.Println("counting")
	// pushed into stack the popped in LIFO manner
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
