package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	var sum int

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: please enter a number.")
		return
	}

	max, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q is not a number\n", args[0])
		return
	} else if max < 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	for i := 1; i <= max; i++ {
		s := " + "
		if i == max {
			s = " = "
		}
		if i%2 == 0 {
			fmt.Printf("%d%s", i, s)
			sum += i
		}
	}
	fmt.Println(sum)
}
