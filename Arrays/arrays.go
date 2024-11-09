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

/*
In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.

Run codeCopy code
package main
import "fmt"
func main() {
Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.

    var a [5]int
    fmt.Println("emp:", a)
We can set a value at an index using the array[index] = value syntax, and get a value with array[index].

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
The builtin len returns the length of an array.

    fmt.Println("len:", len(a))
Use this syntax to declare and initialize an array in one line.

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
You can also have the compiler count the number of elements for you with ...

    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
If you specify the index with :, the elements in between will be zeroed.

    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
You can create and initialize multi-dimensional arrays at once too.
*/
