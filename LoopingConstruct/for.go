package main

import "fmt"

// for is Go’s only looping construct. Here are some basic types of for loops.
func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}
	// for n := range n
	// range [0,n)
	for n := range 6 {
		fmt.Println(n)
	}
}
