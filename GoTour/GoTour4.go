package main

import (
	"fmt"
	"math"
	"time"
	"io"
	"strings"
)

// fourth lesson
// methods, pointers cnd, interfaces, errors and
// streams (beggining with readers)

type Vertex struct {
	X, Y float64
}

// you can methods on type like this
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type MyFloat float64

// methods can be declared with all custom types as receivers but they
// have to be declared in the same package, thus you can't declare 
// methods on built-in types
func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return MyFloat(-f)
	}
	return f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods can be declared on pointer types using *T where T itself
// cannot be pointer type. this notation allows methods to alter
// the receiver in contrast to methods value notation
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


// interface is a type that can be anything that implements methods
// determind in the interfaces set, those functions have to be
// known during compilation (at least thats what i assume) otherwise
// it'll not compile
type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// example of implementing Stringer interface (method String())
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// empty interfaces may hold any type, they are often used to handle
// unknown types for example ftm.Print takes arguments of interface{}
type empty interface {}

// Errors
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	// method with value type of the receiver can take value type
	// and pointer, the same goes for methods with methods with
	// receiver pointers types, difference is that methods with 
	// pointer types can change variable they are used on, the
	// former can't
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

	var a Abser
	f := MyFloat(-math.Sqrt2)
	va := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &va // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = va

	fmt.Println(a.Abs())

	// WARNING, IMPORTANT!!!
	// interfaces methods can be called on nil values, in comparison 
	// in other langauges following call would trigger null pointer
	// exception
	var i I

	// this call here would trigger runtime error, interfaces can be
	// thought as tuple (value, type) type holds type of value 
	// assigned to interface, if no type has been assigned type and
	// value will both hold nils and that will trigger runtime nil error
	// if nil value of some type has been assigned, the type will hold
	// type of pointer and value will hold nil
	// describe(i)
	// i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	// type assertion
	var in interface{} = "hello"

	s := in.(string)
	fmt.Println(s)

	// it can return two values, as always with assertions, default
	// value of type and boolean, if in holds (string) ok = bool
	s, ok := in.(string)
	fmt.Println(s, ok)

	f, ok := in.(float64)
	fmt.Println(f, ok)

	f = in.(float64) // panic
	fmt.Println(f)

	// type switch
	do(21)
	do("hello")
	do(true)

	// Errors
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Readers
	
}

// type switch
func do(i interface{}) {
	// this switch has keyword type in interface assertion function and
	// this means that it returns type wich is held by i interface tuple
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}