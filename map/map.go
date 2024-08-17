// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// /*
// key -> string
// value -> struct of {Long, Float}
// */
// var mp map[string]int

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])
// 	mp = make(map[string]int) //make() funnction creates the map
// 	mp["first"] = 1
// 	mp["second"] = 2
// 	mp["third"] = 3
// 	fmt.Println("Map length is: ", len(mp))
// 	for i := 0; i < len(mp); i++ {
// 		fmt.Println(mp["first"])
// 	}
// 	// check if key-value is present in map or not
// 	v, ok := mp["Answer"]
// 	fmt.Println("The value:", v, "Present?", ok)
// }

/*
	Implement WordCount. It should return a map of the counts of each “word” in the string s.

The wc.Test function runs a test suite against the provided function and prints success or failure
*/
package main

import (
	"fmt"
	"strings"
)

var mp, mp2 map[string]int

func WordCount(str string) map[string]int {
	mp = make(map[string]int)
	fieldsStr := strings.Fields(str)
	for i := 0; i < len(fieldsStr); i++ {
		mp[fieldsStr[i]] += 1
	}
	return mp
}

func main() {
	s := "Hey how are you Hey how do you do ? I am do ing party"

	mp2 = WordCount(s)
	fmt.Println(mp2)

}
