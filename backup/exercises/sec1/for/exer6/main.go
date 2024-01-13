package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers: Verbose Edition
//
//  By using a loop, sum the numbers between 1 and 10.
//
// HINT
//  1. For printing it as in the expected output, use Print
//     and Printf functions. They don't print a newline
//     automatically (unlike a Println).
//
//  2. Also, you need to use an if statement to prevent
//     printing the last plus sign.
//
// EXPECTED OUTPUT
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
		minMaxErr = "min should be lower than max."
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
