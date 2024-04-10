// Getting user input - Reading from standard input
package main

import (
	"fmt"
)

func main() {
	// Get user input
	fmt.Print("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
