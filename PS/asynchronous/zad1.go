package asynchronous

import (
	"fmt"
	"sync"
	"time"
)



func Zad1() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("Hello World!!")
		time.Sleep(10 * time.Second)
		wg.Done()
	}()
	wg.Wait()
}