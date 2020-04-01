package asynchronous

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// TODO: zmienić konsolę na okienka bo przecież czytanie z konsoli kiedy się z niej wypisuje jest bez sensu xD
func Zad2() {
	workers := make([]chan int, MAXWORKERS)
	var command string
	consoleRead := bufio.NewScanner(os.Stdin)
	for workerId := range workers {
		workers[workerId] = make(chan int)
	}

	for i := 0; i < 10; i++ {
		go writerZad2(10 - i - 1, workers[10 - i - 1])
	}

	for {
		consoleRead.Scan()
		command = consoleRead.Text()
		handleCommand(command, workers)
	}
}

func writerZad2(gInd int, lockerChan chan int) {
	instruction := <- lockerChan
	for r := 'A'; ; {
		if instruction == PAUSE {
			instruction = <- lockerChan
		} else {
			select {
			case instruction = <- lockerChan:
			default:
				if r > 'Z' {
					r = 'A'
				}
				fmt.Printf("%s%d\n", string(r), gInd)
				r++
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func handleCommand(command string, workers []chan int) {
	parts := strings.Split(command, " ")
	if len(parts) < 2 {
		log.Println("not enough information (missing action or thread number)")
		return
	}
	action := parts[0]
	gInd, err := strconv.Atoi(parts[1])

	if err != nil || gInd > 9 || gInd < 0 {
		log.Println("Incorrect goroutine indicator, action will not be performed")
		return
	}

	if strings.EqualFold(action, "R") {
		workers[gInd] <- RESUME
	} else if strings.EqualFold(action, "P") {
		workers[gInd] <- PAUSE
	} else {
		log.Printf("Unsupported command: %s\n", action)
	}
}