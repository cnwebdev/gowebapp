package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

const usage = `
Usage: commands [operator] [text]
Available commands: lower, upper and title
`

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	op, str := args[0], args[1]

	switch op {
	case "lower":
		str = strings.ToLower(str)
	case "upper":
		str = strings.ToUpper(str)
	case "title":
		str = strings.ToLower(str)
		str = strings.Title(str)
	default:
		fmt.Printf("Unknown command: %q\n", op)
		return
	}

	fmt.Printf("your String %s => %s\n", args[1], str)

}
