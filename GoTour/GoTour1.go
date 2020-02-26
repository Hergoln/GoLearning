package main

import(
	"fmt"
	"math/cmplx"
)
// variables, types, constants

/*
 existing types:
 bool
 string
 int  int8  int16  int32  int64
 uint uint8 uint16 uint32 uint64 uintptr
 byte // alias for uint8
 rune // alias for int32
      // represents a Unicode code point
 float32 float64
 complex64 complex128 (yes, complex type like 23 + 3i)
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// const cannot be declared and assigned using ":=" constructor
const Pi = 3.14

func weirdCocat(x, y string) (a, b string){
	a = x + y
	b = y + x
	return
}

// outside functions you can only declare functions and variables this way
// using keywords like var, func
var konst = "huhu"

func main() {

	// this ":="" constructor is available only inside functions
	k := 3

	// a, b := weirdCocat("mister!", "Henlo")
	// fmt.Println(a, b)

	var c, python, java = true, false, "no!"
	var x, y = 3, 4

	fmt.Println(c, python, java, k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// conversions example, in Go values have to be explicitly converted
	// unlike in C
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(x, y, z)
	fmt.Println(i, f, u)
}


