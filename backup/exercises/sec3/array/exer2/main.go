package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Get and Set Array Elements
//
//  1. Use the 01-declare-empty exercise
//  2. Remove everything but the array declarations
//
//  3. Assign your best friends' names to the names array
//
//  4. Assign distances to the closest cities to you to the distance array
//
//  5. Assign arbitrary bytes to the data array
//
//  6. Assign a value to the ratios array
//
//  7. Assign true/false values to the alives arrays
//
//  8. Try to assign to the zero array and observe the error
//
//  9. Now use ordinary loop statements for each array and print them
//      (do not use for range)
//
//  10. Now use for range loop statements for each array and print them
//
//  11. Try assigning different types of values to the arrays, break things,
//     and observe the errors
//
//  12. Remove some of the array assignments and observe the loop outputs
//      (zero values)
//
//
// EXPECTED OUTPUT
//
//  Note: The output can change depending on the values that you've assigned to them, of course.
//        You're free to assign any values.
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================

//
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//   FOR RANGES
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================
//
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
