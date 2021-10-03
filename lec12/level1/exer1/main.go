// Other data structures: Maps and Slices
package main

import "fmt"

func main() {

	var names []string

	var distances []int

	var data []uint8

	var ratios []float64

	var alives []bool

	fmt.Printf("names\t\t: %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances\t: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data\t\t: %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios\t\t: %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives\t\t: %T %d %t\n", alives, len(alives), alives == nil)

}

// ---------------------------------------------------------
// EXERCISE: Declare nratios
//
//  1. Declare the following slices as nil slices:
//
//    1. The names of your friends (names slice)
//
//    2. The distances to locations (distances slice)
//
//    3. A data buffer (data slice)
//
//    4. Currency exchange ratios (ratios slice)
//
//    5. Up/Down status of web servers (alives slice)
//
//
//  2. Print their type, length and whether they're equal to nil value or not.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------
