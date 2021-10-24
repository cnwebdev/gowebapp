package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------

func main() {
	// Change this accordingly to produce the expected outputs
	str := os.Args[0:]

	// fmt.Println(len(str), str)

	if len(str) != 2 {
		fmt.Println("Please enter your age?")
		return
	}

	age, err := strconv.Atoi(str[1])
	if err != nil {
		fmt.Println("Not a valid number")
		return
	}

	if age < 0 {
		fmt.Println("Please enter a positiver interger.")
		return
	}

	var res string

	// Type your if statement here.
	if age > 60 {
		res = "Getting older"
	} else if age > 30 {
		res = "Getting wiswer"
	} else if age > 20 {
		res = "Adulthood"
	} else if age > 10 {
		res = "Yound blood"
	} else {
		res = "Booting up"
	}

	fmt.Println(res)

}
