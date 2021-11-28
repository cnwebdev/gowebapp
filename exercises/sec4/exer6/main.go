package main

import (
	"fmt"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesC := strings.Split(namesA, ", ")

	sort.Strings(namesB)
	sort.Strings(namesC)

	var e string

	if len(namesC) != len(namesB) {
		e = "not"

	}
	/*
		for i, name := range namesB {
			if name != namesC[i] {
				e = "not"
				break
			}
		}
	*/
	fmt.Printf("They are %s equal.\n", e)
}
