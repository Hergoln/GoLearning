package main

import (
	"fmt"
	"sync"
	"time"
)

// lesson 5
// bread and butter of Go, goroutines, channels and little bit of
// web with sync package

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// this function is feeding channel with ints
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// sender function, example of channel with close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// sender function with 2 channels, c channel receives values and
// quit channels awaits for value loop stops when quit channel
// receives value
func fibonacciQuit(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	// goroutines, lightweight go coroutines
	// important:
	// evaluation of function and its parameters happen in current goroutine
	// but execution of function happens in new goroutine
	// goroutines runs in the same address space so access to shared memory
	// have to be synchronized
	go say("world")
	say("hello")

	// channels
	s := []int{7, 2, 8, -9, 4, 0}

	// this code sums one half of slice in each goroutine and
	// returns computed ints to x and y respectively
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// buffered channel, give second parameter in make function
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// fibonacci function closes channel after it run out of values
	// to send
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// for receives value from channel until channel is closed
	for i := range c {
		fmt.Println(i)
	}

	// annonymous function goroutine call
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciQuit(c, quit)

	// example with default select case
	// time.Tick(t) returns channel thats fed in t time interval
	tick := time.Tick(100 * time.Millisecond)
	// time After(d) returns channel that is fed in d time delay
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	// sync
	// SafeCounter methods (Inc and Value) use Mutex to ensure
	// only one goroutine can use them at any given time
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
