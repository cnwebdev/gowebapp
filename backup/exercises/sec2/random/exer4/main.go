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
)

func main() {

	var (
		guess int
		v     string
		err   error
		max   = 5
	)

	// Read user number from command-line
	args := os.Args[1:]
	l := len(args)

	// Verify user input syntax
	// Convert user string to number and verify conversion error
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

	// Generate random number
	rand.Seed(time.Now().UnixNano())

	// Compare user number to computer generated number
	// Print random wins/losses messages for each matching case
	for i := 1; i <= max; i++ {
		cRand := rand.Intn(max + 1)

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
