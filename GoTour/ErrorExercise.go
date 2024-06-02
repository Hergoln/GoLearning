package main

import (
	"fmt"
)

type MyError struct {
	What string
}

func (err *MyError) Error() string {
	return fmt.Sprint("What?: %s", err.What)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &MyError{"Negative number"}
	}
	z := 1.
	zP := 0.
	precision := 0.001
	for ; z - zP > precision || z - zP < -precision;
		zP, z = z, (z - (z*z - x) / (2*z)) {
		}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
