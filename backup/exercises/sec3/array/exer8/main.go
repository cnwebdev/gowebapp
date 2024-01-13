package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {

	// create scientest Multi-array
	var scientist = [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "redium"},
		{"Charles", "Darwin", "fittest"},
	}

	// Printf - Printing the heading table
	fmt.Printf("%-15s %-15s %-15s\n", "First Name", "Last Name", "Nickname")
	fmt.Println(strings.Repeat("=", 45))

	// for - cycle through the scientist array and print out each element
	for i := range scientist {
		n := scientist[i]

		for _, v := range n {
			fmt.Printf("%-16s", v)
		}
		fmt.Println()
	}
}
