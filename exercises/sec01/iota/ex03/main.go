package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Seasons
//
//  Use iota to initialize the season constants.
//
// HINT
//  You can change the order of the constants.
//
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------

func main() {
	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	const (
		Spring = (1 + iota) * 3 // 3
		Summer // (1 + 1) * 3 = 6
		Fall   // (1 + 2) * 3 = 9
		Winter  // (1 + 3) * 3 = 12
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}