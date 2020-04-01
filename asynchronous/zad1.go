package asynchronous

import (
	"fmt"
	"time"
)

func Zad1() {
	errChan := make(chan error)
	// create goroutine (golangs thread)
	go func(err chan error) {
		fmt.Println("Hello World!!")
		time.Sleep(10 * time.Second)
		// pass value to channel
		err <- nil
	}(errChan)

	// waiting for value on channel
	fmt.Print(<- errChan)
}