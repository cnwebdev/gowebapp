package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {

	names := [...]string{"Mark", "Tim", "Hans"}
	distances := [...]int{93, 32, 43, 23, 66}
	data := [...]byte{'A', 'B', 'C', 'D', 'F'}
	ratios := [...]float64{3.14}
	alives := [...]bool{true, false, true, false, true}
	zero := [...]uint8{}

	// Print the elements
	fmt.Println("names")
	fmt.Println(strings.Repeat("=", 20))
	for i, na := range names {
		fmt.Printf("names[%d]: %q\n", i, na)
	}

	fmt.Println()
	fmt.Println("distances")
	fmt.Println(strings.Repeat("=", 20))
	for i, di := range distances {
		fmt.Printf("distances[%d] %d\n", i, di)
	}

	fmt.Println()
	fmt.Println("data")
	fmt.Println(strings.Repeat("=", 20))
	for i, da := range data {
		fmt.Printf("data[%d] %v %d %x\n", i, da, da, da)
	}

	fmt.Println()
	fmt.Println("ratios")
	fmt.Println(strings.Repeat("=", 20))
	for i, ra := range ratios {
		fmt.Printf("distances[%d] %.2f\n", i, ra)
	}

	fmt.Println()
	fmt.Println("alives")
	fmt.Println(strings.Repeat("=", 20))
	for i, al := range alives {
		fmt.Printf("distances[%d] %t\n", i, al)
	}

	fmt.Println()
	fmt.Println("zero")
	fmt.Println(strings.Repeat("=", 20))
	for i, ze := range zero {
		fmt.Printf("distances[%d] %d\n", i, ze)
	}
}
