// Banger code example
// Using strings package
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Reading string from command line
	str := os.Args[1]

	// count the number of bytes and save the value to veriable l
	l := len(str)

	// store the repeated string to veriable r
	r := strings.Repeat("!", l)

	// convert the string come command line to upper case and store in veriable s
	s := strings.ToUpper(str)

	// concatinate the data and print to terminal
	fmt.Println(r+s+r)
}