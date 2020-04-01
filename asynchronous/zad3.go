package asynchronous

import (
	"fmt"
	"time"
)
// TODO: zmienić konsolę na okienka bo przecież czytanie z konsoli kiedy się z niej wypisuje jest bez sensu xD

func Zad3() {
	lock := make(chan int, MAXWORKERS)
	for i := 0; i < 9; i++ {
		go writerZad3(10 - i - 1, lock)
	}
	lock  <- 0
	writerZad3(0, lock)
}

func writerZad3(gInd int, lockerChan chan int) {
	for r := 'A'; ; {
		if r > 'Z' {
			r = 'A'
		}
		// simplest "lock" you can do in Go
		<- lockerChan
		fmt.Printf("%s%d\n", string(r), gInd)
		lockerChan <- 0
		r++
		time.Sleep(1 * time.Second)
	}
}
