package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	var (
		i   = 1
		sum int
	)

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

	for {
		if i > max {
			break
		}

		s := " + "
		if i == max {
			s = " = "
		}
		if i%2 == 0 {
			fmt.Printf("%d%s", i, s)
			sum += i
		}

		i++
	}
	fmt.Println(sum)
}
