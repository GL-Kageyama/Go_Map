package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Map Generation
	abcMap := make(map[int]string)
	fmt.Println(abcMap)

	// Value Assignment
	abcMap[1] = "A"
	abcMap[2] = "B"
	abcMap[3] = "C"
	fmt.Println(abcMap)

	// Generation using literals
	abcdeMap := map[int]string{1:"A", 2:"B", 3:"C", 4:"D", 5:"E"}
	fmt.Println(abcdeMap)

	// Element Reference
	fmt.Println(abcdeMap[1])
	fmt.Println(abcdeMap[5])

	// Number of elements
	fmt.Println(len(abcMap))

	// Delete Elements
	delete(abcMap, 2)
	fmt.Println(abcMap)

	// Output by for
	for i, m := range abcdeMap {
		fmt.Println(strconv.Itoa(i) + " : " + m)
	}
}

// ===========================
//       Output Sample
// ===========================
// ~ $ go build map.go 
// ~ $ ./map
// map[]
// map[1:A 2:B 3:C]
// map[1:A 2:B 3:C 4:D 5:E]
// A
// E
// 3
// map[1:A 3:C]
// 1 : A
// 2 : B
// 3 : C
// 4 : D
// 5 : E