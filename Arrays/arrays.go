// The type [n]T is an array of n values of type T.
package main

import (
	"fmt"
)

func main() {
	str := [2]string{"harsh", "business"}
	fmt.Println(str[0], str[1])
	var numArr [5]int // array declaration
	for i := 0; i < 5; i++ {
		numArr[i] = i
	}
	for i := 0; i < 5; i++ {
		fmt.Print(numArr[i])
	}
	fmt.Println('\n')
	arr := [5]int{1, 2, 3, 4, 5}
	var slicedArr []int = arr[2:4] // index:[2,4)
	fmt.Printf("%d \n %d\n", len(slicedArr), cap(slicedArr))
	//The capacity of the slice is the number of elements between
	// the start of the slice (index 2 in this case) and the end of the underlying array (arr). i.e. 5-2 = 3
	fmt.Println(slicedArr)
	/*
		Slices are like references to the array
		Changing the elements of a slice modifies the corresponding elements of its underlying array.
	*/

	// Slice Literals : it is like an array literal without length
	s := []struct {
		i int
		b bool
		c string
	}{
		{2, true, "hey"},
		{3, false, "buddy"},
		{5, true, "lightweight"},
		{7, true, "yeah"},
		{11, false, "bye"},
		{13, true, "byee"},
	}
	fmt.Println(s)

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

	// The make function allocates a zeroed array and returns a slice that refers to that array:

	a := make([]int, 5) // len(a)=5
	//	To specify a capacity, pass a third argument to make:

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
	fmt.Printf("%d  %d", len(a), cap(a))
	fmt.Println()
	fmt.Printf("%d  %d", len(b), cap(b))

}
