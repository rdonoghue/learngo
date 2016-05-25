package main

import (
    "fmt"
)

func swap(x,y string) (string, string) {
  return y,x
}

func main(){
  fmt.Println(swap("Fred","Blue"))

  a, b := swap("Christie", "Purple")
  fmt.Println(a,b)

}
