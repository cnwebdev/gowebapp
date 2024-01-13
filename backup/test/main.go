// Testing the for loop
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Reading input from user (command [operator] [table size])
	args := os.Args[1:]

	// Verify input syntax
	if len(args) != 2 {
		fmt.Println("Please anter command [math operator] [table size]")
		fmt.Println(`Example: command "%*/+-" 5`)
		return
	}

	// Declare variable to hold the operator strings
	var (
		validOp   = "%*/+-"
		res       int
		inValidOp = -1
	)

	// store operator string from command line to op variable
	op := args[0]

	// Verify valid operator
	if (strings.IndexAny(validOp, op)) == inValidOp {
		fmt.Printf("%q is not a valid operator\n", op)
		return
	}

	// store the table size value to tSize variable
	tSize, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%q is not a valid number\n", args[1])
		return
	}

	fmt.Printf("%5s", op)
	for i := 0; i <= tSize; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= tSize; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= tSize; j++ {
			switch op {
			case "*":
				res = i * j
			case "/":
				if j != 0 {
					res = i / j
				}
			case "%":
				if j != 0 {
					res = i % j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}
			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}
