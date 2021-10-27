// if, else if and else statements
package main

import "fmt"

func main() {

	var (
		sam   = 20
		mark  = 23
		older bool
	)

	if sam > mark {
		older = true
	} else {
		older = false
	}

	fmt.Println("Sam is older than mark?", older)

}
