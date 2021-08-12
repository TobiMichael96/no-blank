package main

import (
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

// Main function
func main() {
	var sleepTimer float64
	flag.Float64Var(&sleepTimer, "time", 180, "Set the no-blank value in seconds.")
	flag.Parse()
	tempX, tempY, x, y := 0, 0, 0, 0
	x, y = getMousePosition()
	lastAction := getCurrentTime()
	away := false
	fmt.Println(fmt.Sprintf("No-Blank value set to %.f seconds.", sleepTimer))
	for {
		time.Sleep(1 * time.Second)
		tempX, tempY = getMousePosition()
		if x != tempX && y != tempY {
			if away == true {
				message := generateAwayTime(getTimeDiff(lastAction))
				fmt.Println(message)
				err := beeep.Notify("No-Blank", message, "assets/information.png")
				if err != nil {
					return
				}
			}
			away = false
			x, y = getMousePosition()
			lastAction = getCurrentTime()
		} else {
			if getTimeDiff(lastAction).Seconds() > sleepTimer {
				if away != true {
					message := "You have been away for too long."
					fmt.Println(message)
					err := beeep.Notify("No-Blank", message, "assets/information.png")
					if err != nil {
						return
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

func generateAwayTime(awayTime time.Duration) string {
	if awayTime.Minutes() < 60 {
		return fmt.Sprintf("You have been away for %.1f minute(s).", awayTime.Minutes())
	} else {
		return fmt.Sprintf("You have been away for %.1f hours(s).", awayTime.Hours())
	}
}
