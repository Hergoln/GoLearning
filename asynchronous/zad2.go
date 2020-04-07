package asynchronous

import (
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"time"
)

func Zad2() {
	workers := make([]chan int, MAXWORKERS)
	for workerId := range workers {
		workers[workerId] = make(chan int)
	}

	for i := 0; i < 10; i++ {
		go writerZad2(10 - i - 1, workers[10 - i - 1])
	}

	// I am sorry for this monstrosity but something is wrong with function Append for Boxes in fyne.io lib
	// and I could not use it as intended in a loop
	buttonsResume := widget.NewVBox(widget.NewButton(fmt.Sprintf("Start %d", 0),  func() {
		workers[0] <- RESUME
	}),
		widget.NewButton(fmt.Sprintf("Start %d", 1),  func() {
			workers[1] <- RESUME
		}),
		widget.NewButton(fmt.Sprintf("Start %d", 2),  func() {
			workers[2] <- RESUME
		}),
		widget.NewButton(fmt.Sprintf("Start %d", 3),  func() {
			workers[3] <- RESUME
		}),
		widget.NewButton(fmt.Sprintf("Start %d", 4),  func() {
			workers[4] <- RESUME
		}),
		widget.NewButton(fmt.Sprintf("Start %d", 5),  func() {
			workers[5] <- RESUME
		}),
		widget.NewButton(fmt.Sprintf("Start %d", 6),  func() {
			workers[6] <- RESUME
		}),
		widget.NewButton(fmt.Sprintf("Start %d", 7),  func() {
			workers[7] <- RESUME
		}),
		widget.NewButton(fmt.Sprintf("Start %d", 8),  func() {
			workers[8] <- RESUME
		}),
		widget.NewButton(fmt.Sprintf("Start %d", 9),  func() {
			workers[9] <- RESUME
		}))

	buttonsPause := widget.NewVBox(widget.NewButton(fmt.Sprintf("Stop %d", 0),  func() {
			workers[0] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 1),  func() {
			workers[1] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 2),  func() {
			workers[2] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 3),  func() {
			workers[3] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 4),  func() {
			workers[4] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 5),  func() {
			workers[5] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 6),  func() {
			workers[6] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 7),  func() {
			workers[7] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 8),  func() {
			workers[8] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 9),  func() {
			workers[9] <- PAUSE
		}))

	a := app.New()
	win := a.NewWindow("zad 2")
	win.SetContent(
		widget.NewVBox(
			widget.NewHBox(
				buttonsResume,
				buttonsPause,
			),
			widget.NewButton("Quit", func() {
				a.Quit()
			}),
		))
	win.ShowAndRun()
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