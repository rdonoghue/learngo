package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4
  var a float64 = math.Sqrt(float64(x*x + y*y))
  b := float64(a)
  c := uint(y)
  fmt.Println(a,b,c)
}
