package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------
const (
	age17Plus = "R-Rated"
	age13To17 = "PG-13"
	age13To01 = "PG-Rated"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter a number for age.")
		return
	}

	if age, err := strconv.ParseInt(args[0], 10, 32); err != nil {
		fmt.Printf("Error: %q is not a number\n", args[0])
	} else if age < 0 {
		fmt.Println("Please enter a positive number")
	} else if age > 17 {
		fmt.Println(age17Plus)
	} else if age >= 13 && age <= 17 {
		fmt.Println(age13To17)
	} else {
		fmt.Println(age13To01)
	}
}
