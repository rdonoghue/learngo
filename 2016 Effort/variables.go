package main

import "fmt"

var i, j int = 1, 2

/* check it out - implicit types! */

/* Ok, important to note that "var" allows an "="
but without the var I can implictly assign with :=
so
*/

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	d := false
	e := true
	f := "yes!"
	fmt.Println(i, j, d, e, f)
}
