package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	// Clear the screen
	screen.Clear()

	for {
		// Move cursor to the top left cornor of the screen
		screen.MoveTopLeft()

		// Getting the current time
		now := time.Now()
		hour := now.Hour()
		minute := now.Minute()
		second := now.Second()
		// fmt.Println(hour, minute, second)

		clock := [...]Holder{
			Numbers[hour/10], Numbers[hour%10],
			sep,
			Numbers[minute/10], Numbers[minute%10],
			sep,
			Numbers[second/10], Numbers[second%10],
		}

		for l := range clock[0] {
			for n, d := range clock {
				s := clock[n][l]
				if d == sep && second%2 == 0 {
					s = "   "
				}
				fmt.Print(s, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(100 * time.Millisecond)
	}
}
