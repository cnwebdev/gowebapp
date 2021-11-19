package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments is not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
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

	// Bobble-sort
	for range nums {
		for i, v := range nums {
			if i < len(nums)-1 && v > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	// Print the given numbers and their average
	fmt.Printf("%.2f\n", nums)
	fmt.Printf("Average: %.2f\n", sum/avg)
}
