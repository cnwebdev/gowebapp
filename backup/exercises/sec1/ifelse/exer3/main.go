package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {

	str := os.Args[1:]

	l := len(str)
	// myStr := str[1:]

	fmt.Println(len(str), str)

	if l == 0 {
		fmt.Println("Give me an argument")
		return
	}

	if l == 1 {
		fmt.Printf("There is one: %q\n", str[0])
	} else if l == 2 {
		fmt.Printf("There are two: %q %q\n", str[0], str[1])
	} else {
		fmt.Printf("There are %d %s\n", l, str[0:])
	}

}
