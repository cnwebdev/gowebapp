package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the Temperature
//
//  Print the current temperature in your area using Printf
//
// NOTE
//  Do not use %v verb
//  Output "shouldn't" be like 29.500000
//
// EXPECTED OUTPUT
//  Temperature is 29.5 degrees.
// ---------------------------------------------------------

func main() {

	temp := 67.3

	fmt.Printf("Temperature in Tracy, CA Wed Oct 6, at 9:35 AM is %.1f\n", temp)
}
