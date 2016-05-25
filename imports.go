package main

import (
  "fmt"
  "math"
)

/* note Printf, not Println - that's why I needed to add the /n */

func main() {
  fmt.Printf("Now you have %g problems\n", math.Sqrt(9))
}
