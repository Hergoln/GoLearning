package asynchronous

import (
	"fmt"
	"time"
)

func Zad4() {
	buffer := make(chan string, MAXWORKERS)

	for i := 0; i < 10; i++ {
		go writerZad4(10 - i - 1, buffer)
	}

	for {
		fmt.Println(<- buffer)
	}
}

func writerZad4(gInd int, lockerChan chan string) {
	for r := 'A'; ; {
		if r > 'Z' {
			r = 'A'
		}
		lockerChan <- fmt.Sprintf("%s%d", string(r), gInd)
		r++
		time.Sleep(1 * time.Second)
	}
}