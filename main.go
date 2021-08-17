package main

import (
	"flag"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/beeep"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

// Main function
func main() {
	var sleepTimer float64
	var notification bool
	flag.Float64Var(&sleepTimer, "time", 180, "Set the no-blank value in seconds.")
	flag.BoolVar(&notification, "notification", false, "Enable host notification.")
	flag.Parse()
	fmt.Println(fmt.Sprintf("No-Blank value set to %.f seconds.", sleepTimer))
	fmt.Println(fmt.Sprintf("Notification enabled: %t", notification))

	a := app.New()
	w := a.NewWindow("No-Blank")

	label := widget.NewLabel(fmt.Sprintf("No-Blank time: %0.f seconds", sleepTimer))
	awayTime := widget.NewLabel("Away time: 0 minute(s)")
	progressbar := widget.ProgressBar{Max: sleepTimer / 2}
	w.SetContent(container.NewVBox(
		label,
		&progressbar,
		awayTime,
	))

	go func() {
		tempX, tempY, x, y := 0, 0, 0, 0
		x, y = getMousePosition()
		lastAction := getCurrentTime()
		away := false
		for {
			time.Sleep(1 * time.Second)
			tempX, tempY = getMousePosition()
			if x != tempX && y != tempY {
				progressbar.SetValue(0)
				if away == true {
					message := generateAwayTime(getTimeDiff(lastAction), false)
					awayTime.SetText(generateAwayTime(getTimeDiff(lastAction), true))
					fmt.Println(message)
					if notification {
						err := beeep.Notify("No-Blank", message, "assets/information.png")
						if err != nil {
							return
						}
					}
					time.Sleep(3 * time.Second)
				}
				w.Hide()
				away = false
				x, y = getMousePosition()
				lastAction = getCurrentTime()
			} else {
				if getTimeDiff(lastAction).Seconds() > sleepTimer/2 {
					w.Show()
					awayTime.SetText(generateAwayTime(getTimeDiff(lastAction), true))
					progressbar.SetValue(getTimeDiff(lastAction).Seconds() - sleepTimer/2)
				}
				fmt.Println(getTimeDiff(lastAction).Seconds())
				if getTimeDiff(lastAction).Seconds() > sleepTimer {
					if away != true {
						message := "You have been away for too long."
						fmt.Println(message)
						if notification {
							err := beeep.Notify("No-Blank", message, "assets/information.png")
							if err != nil {
								return
							}
						}
						away = true
					}
					rand.Seed(time.Now().UnixNano())
					robotgo.DragMouse(rand.Intn(1500), rand.Intn(1500))
					time.Sleep(500 * time.Millisecond)
					x, y = getMousePosition()
				}
			}
		}
	}()

	a.Run()
}

func getTimeDiff(offTime time.Time) time.Duration {
	return getCurrentTime().Sub(offTime)
}

func getMousePosition() (int, int) {
	x, y := robotgo.GetMousePos()
	return x, y
}

func getCurrentTime() time.Time {
	return time.Now()
}

func generateAwayTime(awayTime time.Duration, gui bool) string {
	if awayTime.Minutes() < 60 {
		if gui != true {
			return fmt.Sprintf("You have been away for %.1f minute(s).", awayTime.Minutes())
		} else {
			return fmt.Sprintf("Away time: %.1f minute(s)", awayTime.Minutes())
		}
	} else {
		if gui != true {
			return fmt.Sprintf("You have been away for %.1f hours(s).", awayTime.Hours())
		} else {
			return fmt.Sprintf("Away time: %.1f hours(s)", awayTime.Hours())
		}
	}
}
