package main

import (
	"fmt"
	"math/cmplx"
)

/* Wait, check this out - big old var?
Ok, ToBe looks weird because I'm used to those being in a
comma seperates list. That is, this could be:
ToBe bool, MaxInt uint64, z complex128 = false, 1<<64 - 1, cmplx.Sqrt(-5 + 12i)

I think. Will test in a sec.
*/



var (
  ToBe bool, MaxInt uint64, z complex128 = false, 1<<64 - 1, cmplx.Sqrt(-5 + 12i)
/*
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
*/
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
