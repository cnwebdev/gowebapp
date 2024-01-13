// Testing Error hangling
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := os.Args[1]

	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T %d\n", num, num)
}
