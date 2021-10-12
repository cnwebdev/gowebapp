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

	switch true {
	case tnow < 6:
		greet = "Good night!"
	case tnow < 12:
		greet = "Good morning!"
	case tnow < 18:
		greet = "Good afternoon!"
	case tnow < 24:
		greet = "Good evening!"
	default:
		fmt.Println("Please enter time from 1 to 23.")
	}

	fmt.Println(greet)

}
