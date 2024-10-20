package main

import (
	"fmt"
)

	/*
	Check out the type - only declared it once - implicitly
	it applies to all thise before it
	*/

func add(x, y int) int {
	return x + y
}

func main () {
	fmt.Println(add(42, 13))
}
