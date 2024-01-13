package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter a number.")
		return
	}

	if n, err := strconv.ParseInt(args[0], 10, 32); err != nil {
		fmt.Printf("Error: %q is not a number\n", args[0])
	} else if n%8 == 0 {
		fmt.Printf("%d is even and devisible to 8.\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is even number.\n", n)
	} else {
		fmt.Printf("%d is odd number.\n", n)
	}
}
