package asynchronous

import (
	"fmt"
	"time"
)

func Zad3() {
	lock := make(chan int, 1)
	for i := 0; i < 9; i++ {
		go writerZad3(10 - i - 1, lock)
	}
	lock  <- 0
	writerZad3(0, lock)
}

func writerZad3(gInd int, lockerChan chan int) {
	for r := 'A'; ; r++ {
		if r > 'Z' {
			r = 'A'
		}
		// channel receive is blocking operation if used on empty channel
		<- lockerChan
		fmt.Printf("%s%d\n", string(r), gInd)
		// writing to channel is also blocking operation if used on full channel
		lockerChan <- 0
		time.Sleep(1 * time.Second)
	}
}
