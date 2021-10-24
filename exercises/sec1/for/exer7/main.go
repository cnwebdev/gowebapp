package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	const (
		usage = `
	Usage: command [min] [max]
	Example: command 1 5
	`
		minErr    = "%q is not a number.\n"
		maxErr    = "%q is not a number.\n"
		minMaxErr = "min should be smaller than max."
	)

	var (
		min int
		max int
		sum int
		err error
	)

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	min, err = strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf(minErr, args[0])
		return
	}

	max, err = strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf(maxErr, args[1])
		return
	}

	if max < min {
		fmt.Println(minMaxErr)
		return
	}

	for i := min; i <= max; i++ {
		s := " + "
		if i == max {
			s = " = "
		}

		fmt.Printf("%d%s", i, s)
		sum += i
	}
	fmt.Println(sum)
}
