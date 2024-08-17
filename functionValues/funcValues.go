package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// functions are also values i.e functions can be passed within functions and works fine
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	res1 := (hypot(5, 12))
	fmt.Println(res1)
	res2 := compute(hypot)
	fmt.Println(res2)
	res3 := compute(math.Pow)
	fmt.Println(res3)
}
