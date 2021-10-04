package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Raw Concat
//
//  1. Initialize the name variable
//     by getting input from the command line
//
//  2. Use concatenation operator to concatenate it
//     with the raw string literal below
//
// NOTE
//  You should concatenate the name variable in the correct
//  place.
//
// EXPECTED OUTPUT
//  Let's say that you run the program like this:
//   go run main.go inanç
//
//  Then it should output this:
//   hi inanç!
//   how are you?
// ---------------------------------------------------------

func main() {
	// uncomment the code below
	// name := "and get the name from the command-line"

	// replace and concatenate the `name` variable
	// after `hi ` below
	str := os.Args[1:]
	r := " "

	msg := `hi CONCATENATE-NAME-VARIABLE-HERE!
how are you?`

	fmt.Println(msg)

	fmt.Println(`hi ` + str[0] + r + str[1] + r + str[2] + `!
how are you?`)

}