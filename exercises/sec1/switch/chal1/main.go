// Switch statements
// Challenge 1 - Greeting base on time of the day by reading system current time
package main

import (
	"fmt"
	"time"
)

func main() {

	tnow := time.Now().Hour()

	fmt.Printf("the time now is %T %d\n", tnow, tnow)

	var greet string

	switch t := tnow; {
	case t < 6:
		greet = "Good night!"
	case t < 12:
		greet = "Good morning!"
	case t < 18:
		greet = "Good afternoon!"
	case t < 24:
		greet = "Good evening!"
	default:
		fmt.Println("Please enter time from 1 to 23.")
	}

	fmt.Println(greet)

}
