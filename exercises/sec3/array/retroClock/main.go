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
			Sep,
			Numbers[minute/10], Numbers[minute%10],
			Sep,
			Numbers[second/10], Numbers[second%10],
		}

		for l := range clock[0] {
			for next, dn := range clock {
				next := clock[next][l]
				if dn == Sep && second%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
