package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

// Declare wins/losses messages
const (
	usage = `GUESS TO WIN GAME
The computer will randomly generate %d numbers, 
You will win the game if your guess number matches any of 
the computer's generated numbers.

How to play?
Example: command [8]
`
	win1 = `🏆 Wow, you got it the first try!
`
	win2 = `🎉 Congrad, you got me :(!
`
	win3 = `🎖 You got it again!
`
	win4 = `🙌 You are too good, I quit!
`
	loss1 = `👹 Gotcha, you loss!
`
	loss2 = `😭 Oh no, you have to do better!
`
	atoiErr = `%q is not a number!
`
	numErr = `Please enter a positive number!
`
	max = 10
)

func main() {

	// Read user number from command-line
	args := os.Args[1:]
	l := len(args)

	// Verify user input syntax
	if l != 1 {
		fmt.Printf(usage, max)
		return
	}

	// Convert user string to number and verify conversion error
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf(atoiErr, args[0])
		return
	} else if guess < 0 {
		fmt.Println(numErr)
		return
	}

	// Generate random number
	rand.Seed(time.Now().UnixNano())

	// Compare user number to computer generated number
	// Print random wins/losses messages for each matching case
	for i := 0; i < max; i++ {
		cRand := rand.Intn(guess + 1)
		fmt.Println(cRand, guess)
		if cRand == guess {
			switch rand.Intn(5) {
			case 0:
				fmt.Println(win1)
			case 1:
				fmt.Println(win2)
			case 2:
				fmt.Println(win3)
			case 3:
				fmt.Println(win3)
			case 4:
				fmt.Println(win4)
			}
			return
		} else {
			switch rand.Intn(2) {
			case 0:
				fmt.Println(loss1)
			case 1:
				fmt.Println(loss2)
			}
			return
		}
	}
}
