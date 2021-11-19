package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------
const usage = `
Give me a list of maximum 5 numbers.
Example: command [3 4 9 3 2]
`

func main() {

	// Get the list of numbers from command-line
	args := os.Args[1:]

	if len(args) < 1 || len(args) > 5 {
		fmt.Println(usage)
		return
	}

	// Create and array and assign the given numbers to that array
	var (
		sum  float64
		avg  float64
		nums [5]float64
	)

	for i, n := range args {
		num, err := strconv.ParseFloat(n, 64)
		if err != nil {
			continue
		}
		nums[i] = num
		avg++
		sum += num
	}

	// Print the given numbers and their average
	fmt.Printf("%.2f\n", nums)
	fmt.Printf("Average: %.2f\n", sum/avg)
}
