package main

import (
	"fmt"
)

	/* Function: Do X, THEN type
	Not sure if thats intuitive or not
	*/

func add(x int, y int) int {
	return x + y
}

func main () {
	fmt.Println(add(42, 13))
}
