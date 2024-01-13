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
	case "united states":
		fmt.Println("Washington D.C.")
	case "japan":
		fmt.Println("Tokyo")
	case "germany":
		fmt.Println("Berlin")
	case "france":
		fmt.Println("Paris")
	case "vietnam":
		fmt.Println("Ha Noi")
	default:
		fmt.Println("I don't know the answer!")
	}
}
