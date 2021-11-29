package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Append #2
//
//  1. Create the following nil slices:
//     + Pizza toppings
//     + Departure times
//     + Student graduation years
//     + On/off states of lights in a room
//
//  2. Append them some elements (use your creativity!)
//
//  3. Print all the slices
//
//
// EXPECTED OUTPUT
// (Your output may change, mine is like so:)
//
//  pizza       : [pepperoni onions extra cheese]
//
//  graduations : [1998 2005 2018]
//
//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
//  lights      : [true false true]
//
//
// HINTS
//  + For departure times, use the time.Time type. Check its documentation.
//
//      now := time.Now()     -> Gives you the current time
//      now.Add(time.Hour*24) -> Gives you a time.Time 24 hours after `now`
//
//  + For graduation years, you can use the int type
// ---------------------------------------------------------

func main() {

	//
	var pizza []string
	var departure []time.Time
	var graduation []int

	pizza = append(pizza, "Pepperoni", "onion", "extra cheese")

	now := time.Now()

	departure = append(departure,
		now,
		now.Add(time.Hour*24),
		now.Add(time.Hour*48))

	graduation = append(graduation, 1998, 2005, 2018)

	fmt.Printf("%T %d %t\n", pizza, len(pizza), pizza == nil)
	fmt.Printf("%T %d %t\n", departure, len(departure), departure == nil)
	fmt.Printf("%T %d %t\n", graduation, len(graduation), graduation == nil)

	fmt.Printf("%T %v\n", pizza, pizza)
	fmt.Printf("%T %v\n", departure, departure)
	fmt.Printf("%T %v\n", graduation, graduation)
}
