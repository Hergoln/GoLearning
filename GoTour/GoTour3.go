package main

import (
	"fmt"
	"strings"
)

// third lesson
// pointers, structs, arrays, slices, for : range, maps, functions

type Vertex struct {
	X int
	Y int
}

// struct literals
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	pl  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	i, j := 42, 2701

	// unlike C, Go has no pointer arithmetic
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// structs
	// construct struct
	v := Vertex{1, 2}
	v.X = 4

	// assign v addres to pointer pv of type *Vertex
	pv := &v 
	pv.X = 1000 // same as (*pv).X

	// basics of arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slice
	// the difference is that arrays have defined size [n]T
	// where slices are flexible and do not contain size []T
	// a[low : high] includes low and excludes high ends of sliced array
	var s []int = primes[1:4]
	fmt.Println(s)

	// slices are "references to part of an array", this example
	// ilustrates how change in slice corresponds with change in array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	an := names[0:2]
	bn := names[1:3]
	fmt.Println(an, bn)

	bn[0] = "XXX"
	fmt.Println(an, bn)
	fmt.Println(names)

	// slices might also be declared as arrays with their length
	// being determind based on number of values during initialization
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	ex := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ex)

	// while making slice omitting one of range bounds they are maxed
	// by default
	exslice := q[:] // is equivalent to q[0:6]

	// slices default value is nill, its length and capacity are 0
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))

	// create slice using built-in make function in order to create
	// dynamically-sized array
	sa := make([]int, 5)
	printSlice(sa)

	sb := make([]int, 0, 5)
	printSlice(sb)

	sc := sb[:2]
	printSlice(sc)

	sd := sc[2:5]
	printSlice(sd)

	// slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}


	// appending to slices
	// append function return slice, if appending to slice/array
	// exceeds capacity of original array, new array will be 
	// allocated and slice of this array will be returned
	var apps []int
	printSlice(apps)

	// append works on nil slices.
	apps = append(apps, 0)
	printSlice(apps)

	// The slice grows as needed.
	apps = append(apps, 1)
	printSlice(apps)

	// We can add more than one element at a time.
	apps = append(apps, 2, 3, 4)
	printSlice(apps)

	// for range example, primes is an array
	// i is index and v is value under that index
	for i, v := range primes {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// this variation returns only indexes
	for i := range primes {
		primes[i] = 1 << uint(i) // == 2**i
	}
	// this variation returns only values
	for _, value := range primes {
		fmt.Printf("%d\n", value)
	}

	// maps
	// map default value is nil, maps works almost the same as
	// arrays except their index might be of different type tham
	// int
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, -74,
	}
	fmt.Println(m["Bell Labs"])

	// map literal with static size
	var mbis = map[string]Vertex{
		"Bell Labs": Vertex{
			40, -74,
		},
		"Google": Vertex{
			37, -122,
		},
	}

	// ^ v those two are equivalent
	var mter = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	// basic operations on maps
	mop := make(map[string]int)

	mop["Answer"] = 42
	fmt.Println("The value:", mop["Answer"])

	mop["Answer"] = 48
	fmt.Println("The value:", mop["Answer"])

	delete(mop, "Answer")
	fmt.Println("The value:", mop["Answer"])

	v, ok := mop["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// functions can also be created local variables like that
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// using closures 
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// printing slices length and capacity
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// functions are also values and they can be passed as arguments
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Go functions may be closures and that means they can reference 
// values from outside their bodies but variables and arguments
// within their bodies are exclusive to given call of closure 
// function
// this example uses sum as accumulator
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}