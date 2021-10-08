// else if example
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	myNum := 9

	yourGuest := os.Args[1]
	fmt.Println(len(yourGuest), yourGuest)

	yGuest, err := strconv.Atoi(yourGuest)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result string

	if yGuest == myNum {
		result = "you won!"
	} else if yGuest > myNum {
		result = "try lower"
	} else if yGuest < myNum {
		result = "try higher"
	} else {
		result = "Something went wrong!"
	}

	fmt.Println(result)

}
