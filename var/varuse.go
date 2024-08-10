package main

import (
	"fmt"
)

// var makes the variables a list of variables and type at last(optional)
func main() {
	var city, day, boy = "Dehradun", 10, false

	/*
		Short variable declarations
		Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

		Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	*/
	yourName, age := "Narendra", 21
	fmt.Println(city, day, boy, yourName, age)
}
