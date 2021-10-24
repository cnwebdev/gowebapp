// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Declare wins/losses messages
const (
	usage = `GUESS TO WIN GAME
The computer will randomly generate %d numbers, 
You will win the game if your guess number matches any of 
the computer's generated numbers.
	
How to play?
You can enter upto 2 guesses each round.

Example 1: command [8]
Example 2: command [8] [9]
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
	numErr = `Please enter positive numbers!
	`
)

func main() {
	var (
		guess1, guess2 int
		err            error
		max            = 10
	)
	// Read user number from command-line
	args := os.Args[1:]
	l := len(args)

	// Verify user input syntax
	if l < 1 {
		fmt.Printf(usage+"\n", max)
		return
	}

	// Convert user string to number and verify conversion error
	guess1, err = strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf(atoiErr, args[0])
		return
	}

	if l == 2 {
		guess2, err = strconv.Atoi((args[1]))
		if err != nil {
			fmt.Printf(atoiErr, args[1])
		}
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println(numErr)
		return
	}

	// Generate random number
	rand.Seed(time.Now().UnixNano())

	// Compare user number to computer generated number
	// Print random wins/losses messages for each matching case
	for i := 0; i < max; i++ {
		var cRand int
		if guess1 < guess2 {
			cRand = rand.Intn(guess2 + 1)
		} else {
			cRand = rand.Intn(guess1 + 1)
		}

		fmt.Println(cRand, guess1, guess2)
		if cRand == guess1 || cRand == guess2 {
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
