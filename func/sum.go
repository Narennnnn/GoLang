package main

import (
	"fmt"
)

// you can notice that return type is used after variable declaration
//
//	func sum(x int, y int) int {
//		return x + y
//	}
func sum(x, y int) int { //if sharing same type then only last variable is needed to declared
	return x + y
}
func main() {
	fmt.Println(sum(100, 101))
}
