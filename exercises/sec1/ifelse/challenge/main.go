// Challenge, adding custom error hangling message to the program
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var meter float64

	str := os.Args[1:]

	if len(str) != 1 {
		fmt.Println("Please enter a number")
		return
	}

	if feet, err := strconv.ParseFloat(str[0], 64); err != nil {
		fmt.Printf("%q is not a number.\n", str)
	} else {
		meter = feet * 0.3048
		fmt.Printf("%.2f feet(s) = %.2f meters\n", feet, meter)
	}
	// feet := float64(n)

}
