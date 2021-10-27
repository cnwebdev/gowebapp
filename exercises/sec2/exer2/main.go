package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Path Searcher
//
//  Your program should search inside the path environment
//  variable.
//
//  Remove the corpus constant then get the corpus from the
//  environment variable "Path" or "PATH".
//
// HINTS
//  1. Search the web to find out what is an environment
//     variable and how to use it (if you don't know
//     what it is already).
//
//  2. Look up for the necessary functions for getting
//     an environment variable. It's in the "os" package.
//
//     Search for it on the Go online documentation.
//
//  3. Look up for the necessary function for splitting
//     the path variable into directory names.
//
//     It's in the "path/filepath" package.
//
// EXAMPLE
//  For example, on my Mac, my PATH environment variable
//  looks like this:
//
//    "/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//
//    go run main.go /sbin
//
//  It should print this:
//
//    #2 : "/sbin"
// ---------------------------------------------------------

// ---------------------------------------------------------
// BONUS EXERCISE
//  Make your program cross platform. So, it can search
//  the path environment variable when you run it on
//  a Windows or on a Mac (OS X) or on a Linux.
//
// BONUS EXERCISE #2
//  Also find the directories for the directories that
//  contains the searched query.
//
//  And let it match in a case-insensitive manner. See the examples.
//
//  EXAMPLE:
//    Let's say:
//      PATH="/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//    go run main.go bin
//
//  It should print this:
//    #1 : "/usr/local/bin"
//    #2 : "/sbin"
//    #3 : "/Users/inanc/go/bin"
//
//  Or like this (case insensitive):
//    go run main.go Us
//
//  Output:
//    #1 : "/usr/local/bin"
//    #2 : "/Users/inanc/go/bin"
// ---------------------------------------------------------

const (
	usage = `Path Search
The program reads user queries and read the computer local environment veriables PATH,
then print out the matching search result.

example: command [/Users/name/go]
`
)

func main() {
	// os.Getenv returns a string with environment directory path.
	envPath := os.Getenv("PATH")
	fmt.Println(envPath)

	// Alternative function
	// test := filepath.SplitList(envPath)
	// fmt.Println(test)

	// strings.Split return a []string with environment path
	path := strings.Split(envPath, ":")
	fmt.Println(path)

	// Read the user's query string from command-line
	args := os.Args[1:]

	// Verity the input syntax
	if len(args) < 1 {
		fmt.Println(usage)
		return
	}

	query := args
	// query = strings.ToLower(query)

	for _, s := range path {
		for j, q := range query {
			s = strings.ToLower(s)
			q = strings.ToLower(q)
			if strings.Contains(s, q) {
				fmt.Printf("#%-5d %s\n", j+1, s)
			}
		}
	}
}
