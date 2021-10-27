// Getting input from command line terminal
package main

import (
	"fmt"
	"os"
)

func main() {

	// reading command line is os.Args package
	args := os.Args

	// os.Args read the command line arguments and store the data in the slice[]string
	fmt.Printf("%T %#v\n", args, args)
}