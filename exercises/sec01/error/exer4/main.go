package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Solution from previous exercise
	/*
		if len(os.Args) != 2 {
			fmt.Println("Give me a year number")
			return
		}

		year, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("%q is not a valid year.\n", os.Args[1])
			return
		}

		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			fmt.Printf("%d is a leap year.\n", year)
		} else {
			fmt.Printf("%d is not a leap year.\n", year)
		}

	*/

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter a year number.")
		return
	}

	if y, err := strconv.ParseInt(args[0], 10, 64); err != nil {
		fmt.Printf("%q is not a valid year.\n", args[0])
	} else if y%4 == 0 || y%400 == 0 {
		fmt.Printf("%d is leap year.\n", y)
	} else {
		fmt.Printf("%d is not a leap year\n", y)
	}
}

// ---------------------------------------------------------
// EXERCISE: Simplify the Leap Year
//
//  1. Look at the solution of "the previous exercise".
//
//  2. And simplify the code (especially the if statements!).
//
// EXPECTED OUTPUT
//  It's the same as the previous exercise.
// -
