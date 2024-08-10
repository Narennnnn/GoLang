package main

import "fmt"

/*
We can return more than one values
*/
func swap(z, y string) (string, string) {
	return y, z
}

func main() {
	s1, s2 := swap("hey", "harsh")
	fmt.Printf("%s  %s", s1, s2)
}
