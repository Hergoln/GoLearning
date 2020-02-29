package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fst := 0
	snd := 1
	return func() int {
		third := fst + snd
		fst = snd
		snd = third
		return third
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
