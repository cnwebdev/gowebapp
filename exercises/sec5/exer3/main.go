package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Append #3 — Fix the problems
//
//  Fix the problems in the code below.
//
// BONUS
//
//  Simplify the code.
//
// EXPECTED OUTPUT
//
//  toppings: [black olives green peppers onions extra cheese]
//
// ---------------------------------------------------------

func main() {
	toppings := []string{"black olives", "green peppers"}
	toppings = append(toppings, "onions", "extra cheese")

	var pizza []string
	pizza = append(pizza, toppings...)

	fmt.Printf("pizza       : %s\n", pizza)
	fmt.Printf("toppings    : %s\n", pizza)
}
