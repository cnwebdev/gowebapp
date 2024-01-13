package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

const (
	usage = `GUESS TO WIN GAME
The computer will randomly generate %d numbers,
You will win the game if your guess number matches any of
the computer's generated numbers.

How to play?
Example: command [-v] [8]
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
)

func main() {

	var (
		guess int
		v     string
		err   error
		max   = 5
		min   int
	)

	// Read user number from command-line
	args := os.Args[1:]
	l := len(args)

	// Verify user input syntax

	if l == 1 {
		// Convert user string to number and verify conversion error
		guess, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf(atoiErr, args[0])
			return
		} else if guess < 0 {
			fmt.Println(numErr)
			return
		}
	} else if l == 2 {
		v = args[0]
		// Convert user string to number and verify conversion error
		guess, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf(atoiErr, args[1])
			return
		} else if guess < 0 {
			fmt.Println(numErr)
			return
		} else if v != "-v" {
			fmt.Println("Try this: command [-v] [8].")
			return
		}
	} else {
		fmt.Printf(usage, max)
		return
	}

	// Generate random number
	rand.Seed(time.Now().UnixNano())

	// Setup enough random number
	min = 5
	if guess > 10 {
		min = guess + (.05 * 100)
	} else if guess > 30 {
		min = guess + (.07 * 100)
	} else if guess > 50 {
		min = guess + (.1 * 100)
	} else if min < guess {
		min = guess
	}

	fmt.Println(min)

	// Compare user number to computer generated number
	// Print random wins/losses messages for each matching case
	for i := 1; i <= max; i++ {
		cRand := rand.Intn(min + 1)

		if v == "-v" {
			fmt.Printf("%d %d %d \n", i, cRand, guess)
			// fmt.Printf("%d %d %d ", i, cRand, guess)
		}

		if cRand == guess {
			if i == 1 {
				fmt.Println(win1)
				return
			}
			switch rand.Intn(5) {
			case 0:
				fmt.Println(win2)
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
		}
	}
	switch rand.Intn(2) {
	case 0:
		fmt.Println(loss1)
	case 1:
		fmt.Println(loss2)
	}
}
