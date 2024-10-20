package main

import (
  "fmt"
)

func main() {
  for x :=0; x<10; x++ {
    fmt.Print(x)
    if x > 4 {
      fmt.Print(" BING")
    } else if x < 4{
      fmt.Print(" BONG")
    }
    fmt.Println()
  }
}
