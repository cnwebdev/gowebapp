package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the Type #2
//
//  Print the type and value of 3.14 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3.14 is float64
// ---------------------------------------------------------

func main() {
	pi := 3.14

	fmt.Printf("Type of %.2f is %T\n", pi, pi)
}
