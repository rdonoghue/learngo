package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
  /* Whoah.  FILO.  */
  defer fmt.Println("last?")
}
