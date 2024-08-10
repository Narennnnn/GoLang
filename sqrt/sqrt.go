package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = (z - math.Pow(z, 2) - x) / (z * 2)
	}
	return z
}
func main() {
	var numb float64
	//Scan() to take input from the user
	fmt.Scan(&numb)
	fmt.Println(sqrt(numb))
}
