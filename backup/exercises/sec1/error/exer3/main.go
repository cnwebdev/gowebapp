package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter a year number.")
		return
	}

	if y, err := strconv.ParseInt(args[0], 10, 64); err != nil {
		fmt.Printf("%q is not a valid year.\n", args[0])
	} else if y%4 == 0 || y%400 == 0 {
		fmt.Printf("%d is leap year.\n", y)
	} else {
		fmt.Printf("%d is not a leap year\n", y)
	}
}
