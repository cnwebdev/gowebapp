package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {

	names := [3]string{"Mark", "Tim", "Hans"}
	distances := [5]int{93, 32, 43, 23, 66}
	data := [5]byte{'A', 'B', 'C', 'D', 'F'}
	ratios := [1]float64{3.14}
	alives := [5]bool{true, false, true, false, true}
	zero := [0]uint8{}

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
