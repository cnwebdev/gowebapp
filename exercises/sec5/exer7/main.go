package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"

	//   1. Convert the string to an []int
	ds := strings.Split(data, " ")

	var nums []int

	for _, s := range ds {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	//   2. Print the slice
	fmt.Println("nums:", nums)

	//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
	evenSlice := nums[0:3]
	fmt.Println("evenSlice:", evenSlice)

	//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
	oddSlice := nums[3:]
	fmt.Println("oddSlice:", oddSlice)

	//   5. Slice it for the two numbers at the middle
	middle := nums[2:4]
	fmt.Println("middle:", middle)

	//   6. Slice it for the first two numbers
	firstTwo := nums[0:2]
	fmt.Println("firstTwo:", firstTwo)

	//   7. Slice it for the last two numbers (use the len function)
	l := len(nums)
	lastTwo := nums[l-2:]
	fmt.Println("lastTwo:", lastTwo)

	//   8. Slice the evens slice for the last one number
	evenLastTwo := evenSlice[2:]
	fmt.Println("evenLastTwo:", evenLastTwo)

	//   9. Slice the odds slice for the last two numbers
	oddLastTwo := oddSlice[1:]
	fmt.Println("oddLastTwo:", oddLastTwo)
}
