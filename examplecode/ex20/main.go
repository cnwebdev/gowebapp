// Loop in go
package main

import "fmt"

func main() {

	var (
		p, s int
		l    = 100
	)

	// Loop and multiply 1 = 1000
	for i := 1; i <= 100; i++ {
		p = i * i
	}
	// Print the result
	fmt.Printf("The product of 1 to 100 is %v\n", p)

	// Loop and add 1 to 1000
	for a := 1; a <= 100; a++ {
		s += a
	}
	fmt.Printf("The sum of 1 to 100 is %d\n", s)

	// Loop and subtract 1000 to 1
	for j := 100; j >= 1; j-- {
		l -= 1
	}
	fmt.Printf("(100 - 1) 100 times is %d\n", l)
}
