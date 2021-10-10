// Scope of variables
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// args, and n variables are only accessable inside the if/else block
	if args := os.Args[1:]; len(args) != 1 {
		fmt.Println("Give me a number")
	} else if n, err := strconv.Atoi((args[0])); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s * 5 = %d\n", args[0], n*5)
	}
}
