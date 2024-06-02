package asynchronous

import (
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"time"
)

func Zad4() {
	buffer := make(chan string, MAXWORKERS)
	closeChan := make([]chan interface{}, 11)
	for i := range closeChan {
		closeChan[i] = make(chan interface{})
	}

	for i := 0; i < 10; i++ {
		go writerZad4(i, buffer, closeChan[i])
	}

	go func () {
		for {
			select {
			case  <- closeChan[10]:
				for i := range closeChan {
					closeChan[i] <- PAUSE
				}
				close(buffer)
				return
			default:
				fmt.Println(<- buffer)
			}
		}
	}()

	a := app.New()

	// I am sorry for this monstrosity but something is wrong with function Append for Boxes in fyne.io lib
	// and I could not use it as intended in a loop
	buttonsPause := widget.NewHBox(widget.NewButton(fmt.Sprintf("Stop %d", 0),  func() {
			closeChan[0] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 1),  func() {
			closeChan[1] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 2),  func() {
			closeChan[2] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 3),  func() {
			closeChan[3] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 4),  func() {
			closeChan[4] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 5),  func() {
			closeChan[5] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 6),  func() {
			closeChan[6] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 7),  func() {
			closeChan[7] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 8),  func() {
			closeChan[8] <- PAUSE
		}),
		widget.NewButton(fmt.Sprintf("Stop %d", 9),  func() {
			closeChan[9] <- PAUSE
		}),
		widget.NewButton("Quit",  func() {
			closeChan[10] <- PAUSE
			a.Quit()
		}),)

	win := a.NewWindow("zad 4")
	win.SetContent(buttonsPause)
	win.ShowAndRun()
}

func writerZad4(gInd int, bufferChan chan string, closeChan chan interface{}) {
	for r := 'A'; ; r++ {
		select {
			case <- closeChan:
				return
			default:
				if r > 'Z' {
					r = 'A'
				}
				bufferChan <- fmt.Sprintf("%s%d", string(r), gInd)
				time.Sleep(1 * time.Second)
		}
	}
}