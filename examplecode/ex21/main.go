// Creat multiplication table
package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		p int
		t = 10
	)

	// Print a new line for better output view
	fmt.Println()

	// Print horizontal header X 0 1 2 3 4 5...10 (5 spaces a part)
	// Print "X" after 5 blank spaces
	fmt.Printf("%5s %5s", "X", "|")
	for i := 0; i <= t; i++ {
		// Print 0 to 10 5 spaces a part
		fmt.Printf("%5d", i)
	}
	// Print new line after 10
	fmt.Println()

	// Print horizontal line
	fmt.Println(strings.Repeat("_", 67))
	fmt.Println()

	// Print vertical header 0 1 2 3...10
	for i := 0; i <= t; i++ {
		fmt.Printf("%5d %5s", i, "|")
		for j := 0; j <= t; j++ {
			p = i * j
			fmt.Printf("%5d", p)
		}
		fmt.Println()
	}
	// Print a new line for better output view
	fmt.Println()

}
