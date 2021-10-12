// Multiple case condition
// Switch Statements
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter a country name?")
		return
	}

	country := args[0]
	country = strings.ToLower(country)

	switch country {
	case "united states", "canada", "mexico":
		fmt.Println("Northern America")
	case "japan", "china", "viet nam":
		fmt.Println("Asia")
	case "germany", "france", "austria":
		fmt.Println("Europe")
	case "brazil", "colombia", "peru":
		fmt.Println("South America")
	default:
		fmt.Println("I don't know the answer!")
	}
}
