package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	city   string     = "billu badmosh"
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", city, city)
	i := 99
	ii := float64(i)
	fmt.Printf("%g \n", ii/2)
	// Constants cannot be declared using the := syntax.
	const pi = 2.99
	fmt.Printf("%g \n", pi)
}

/*
default values:
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/
