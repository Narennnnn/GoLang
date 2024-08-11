package main

import (
	"fmt"
)

func main() {
	i, j := 100, 1001
	p := &i
	fmt.Println(*p) //value at the address pointing to i
	p = &j
	fmt.Println(p) // memory address
}
