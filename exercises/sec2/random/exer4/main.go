package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
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
	max = 5
)

func main() {

	var (
		guess int
		v     string
		err   error
	)

	// Read user number from command-line
	args := os.Args[1:]
	l := len(args)

	// Verify user input syntax

	if l == 1 {
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

	// Convert user string to number and verify conversion error

	// Generate random number
	rand.Seed(time.Now().UnixNano())

	// Compare user number to computer generated number
	// Print random wins/losses messages for each matching case
	for i := 0; i < max; i++ {
		cRand := rand.Intn(guess + 1)

		if v == "-v" {
			fmt.Printf("%d %d %d \n", i, cRand, guess)
			// fmt.Printf("%d %d %d ", i, cRand, guess)
		}

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
		}
	}
	switch rand.Intn(2) {
	case 0:
		fmt.Println(loss1)
	case 1:
		fmt.Println(loss2)
	}
}

////////////////////////////////////////////////////////////////////////////
/*
const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
(Provide -v flag to see the picked numbers.)
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var verbose bool

	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if verbose {
			//fmt.Printf("%d ", n)
			fmt.Printf("%d %d %d "+"\n", turn, n, guess)
		}

		if n == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
*/
