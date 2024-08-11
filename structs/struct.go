package main

import (
	"fmt"
)

type Dimensions struct {
	X int
	Y int
}

func main() {
	obj := Dimensions{100, 10}
	obj.X = 99
	fmt.Println(obj.X, obj.Y)
	ptr := &obj // pointer to the struct
	ptr.X = 11
	fmt.Println(ptr.X)
	obj2 := Dimensions{X: 29}
	fmt.Println(obj2.Y) //by default 0
}
