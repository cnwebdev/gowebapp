// Creat multiplication table base on command line input
package main

// ---------------------------------------------------------
// EXERCISE: Dynamic Table
//
//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 table
//
// RESTRICTION
//  Solve this exercise without looking at the original
//  multiplication table exercise.
//
// HINT
//  There was a max constant in the original program.
//  That determines the size of the table.
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Give me the size of the table
//
//  go run main.go -5
//    Wrong size
//
//  go run main.go ABC
//    Wrong size
//
//  go run main.go 1
//    X    0    1
//    0    0    0
//    1    0    1
//
//  go run main.go 2
//    X    0    1    2
//    0    0    0    0
//    1    0    1    2
//    2    0    2    4
//
//  go run main.go 3
//    X    0    1    2    3
//    0    0    0    0    0
//    1    0    1    2    3
//    2    0    2    4    6
//    3    0    3    6    9
// ---------------------------------------------------------

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
Usage: command [5]
`

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	t, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s is not a number.\n", args[0])
		return
	} else if t < 0 {
		fmt.Println("I can only work with positive interger.")
		return
	}

	var (
		p int
	)

	// Print a new line for better output view
	fmt.Println()

	// Print horizontal header X 0 1 2 3 4 5...10 (5 spaces a part)
	// Print "X" after 5 blank spaces
	fmt.Printf("%5s %5s", "X", "|")
	for i := 0; i <= t; i++ {
		// Print 0 to 10 5 spaces a part
		fmt.Printf("%5d", i)
	}
	// Print new line after 10
	fmt.Println()

	// Print horizontal line
	fmt.Println(strings.Repeat("_", 70))
	fmt.Println()

	// Print vertical header 0 1 2 3...10
	for i := 0; i <= t; i++ {
		fmt.Printf("%5d %5s", i, "|")
		for j := 0; j <= t; j++ {
			p = i * j
			fmt.Printf("%5d", p)
		}
		fmt.Println()
	}
	// Print a new line for better output view
	fmt.Println()

}
