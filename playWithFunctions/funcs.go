package main

import "fmt"

func testing(myFunc func(x int) int) {
	fmt.Println(myFunc(49))
}

// function closures it is called
func main() {
	actAsVariable := func() {
		fmt.Println("Cool naw!")
	}
	actAsVariable()

	test1 := func(x int) int {
		return x + x
	}
	testing(test1)

	func(x, y int) {
		fmt.Println(x * y)
	}(10, 2)

}
