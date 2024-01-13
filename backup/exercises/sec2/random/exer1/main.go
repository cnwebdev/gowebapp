package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

// Declare program instruction
const (
	usage = `Your Lucky Number
The computer will generate %d random numbers, 
if your input number match any of the computer generated number,
You have won the game, a bonus point if you number match
the computer's first random number.

Want to play?...

Usage: command [8]
`

	first = `CONGRATULATION, YOU HAVE WON THE FIRST PRISE!
`

	standard = `You got a match, good game!
`

	atoiErr = `%q is not a number
`

	numErr = `Please enter a positive number
`
	max = 5
)

func main() {
	// Game instruction message
	// fmt.Printf(game+"\n", max)

	// Read user guess number from command-line
	args := os.Args[1:]

	// Check for command line syntax errors
	if len(args) != 1 {
		fmt.Printf(usage, max)
		return
	}

	// Convert user input string to number
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf(atoiErr, args[0])
		return
	}

	// Check for negative user input number
	if guess < 0 {
		fmt.Println(numErr)
		return
	}

	// Set computer random number
	rand.Seed(time.Now().UnixNano())

	// Loop: generate 5 random number
	for i := 0; i < max; i++ {
		cRand := rand.Intn(guess + 1)

		// compare user number with computer generate number
		if i == 0 && cRand == guess {
			// Display special wining message if user input number mach first computer generated number
			fmt.Println(i, cRand, guess, first)
			return
		} else if cRand == guess {
			// Display winer message other wise
			fmt.Println(i, cRand, guess, standard)
			return
		}
		fmt.Println(i, cRand, guess)
	}
	// Display game loss message
	fmt.Println("You LOSS!, try again?")
}
