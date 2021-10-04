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

	namesASplited := strings.Split(namesA, ", ")
	sort.Strings(namesASplited)

	fmt.Printf("%T %q\n", namesASplited, namesASplited)

	sort.Strings(namesB)
	fmt.Printf("%T %q\n", namesB, namesB)

	if len(namesASplited) == len(namesB) {
		for i := range namesASplited {
			if namesASplited[i] != namesB[i] {
				return
			}
		}
		fmt.Println("nameASplited and namesB are equal.")
	}
}