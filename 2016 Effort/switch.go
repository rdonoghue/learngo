package main

import (
  "fmt"
)

func main() {
  for x :=0; x<10; x++ {
    fmt.Print(x)
    switch x {
    case 1:
      fmt.Print(" foo")
    default:
      fmt.Print(" bar")
      }
    fmt.Println()
    }
  }
