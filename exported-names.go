package main

import (
	"fmt"
	"math"
)

	/*
	point of the lesson is that math.pi
	will fail because the "exported" value
	(ont that comes along with the package)
	needs to be capitalized
	*/

func main() {
	fmt.Println(math.Pi)
}
