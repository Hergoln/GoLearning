package main

import(
	"fmt"
	"math"
	"runtime"
	"time"
)

// second lesson
// loops, conditions, switches and defers


// standard if, ifs and fors must have braces {}
func sqrtIf(x float64) string {
	if x < 0 {
		return sqrtIf(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// in go ifs like in the fors you can start with short statement that
// executes before condition is checked, any variable declared in this
// statement exists only ifs and elses scope
func powIf(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here
	return lim
}


func main() {

	sum := 0
	// standard for structure
	for i := 0; sum < 10; i++ {
		sum += 1
	}

	// C's while is spelled for in Go
	for sum <= 100 {
		sum += sum
	}

	// infinite loop may look like this, just omit everything except for
	// for {
	// }

	fmt.Println(sqrtIf(2), sqrtIf(-4))
	fmt.Println(
		powIf(3, 2, 10),
		powIf(3, 3, 20),
	)

	z, iter := Sqrt(16, .01)
	fmt.Printf("Exercise Sqrt(2) = %g %d", z, iter)

	// important thing is cases have builtin break statement
	// that means only one case is runing 
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch wihout condition is equvalent to switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// keyword defer defers execution of a function until
	// surrounding functions returns, in this case main() function
	// defered functions are pushed onto LIFO stack, thus they are
	// called from last to first
	defer fmt.Println("world")

	fmt.Println("hello")
}

// exercise Sqrt with loop
func Sqrt(x float64, precision float64) (res float64, iter int) {
	z := 1.
	zP := 0.
	for ; z - zP > precision || z - zP < -precision;
		zP, z, iter = z, (z - (z*z - x) / (2*z)), iter + 1 {
			fmt.Println(z, zP)
		}
	return z, iter
}