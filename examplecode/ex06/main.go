// Convert value with type conversion
package main

import "fmt"

func main() {

	// Short declaration when initial value is known
	apple := 50
	rice := 5
	iceScream := 2
	
	applePrice := 1.20
	ricePrice := 2.16
	icePrice := 5.60

	// Grouping variable for similar operation
	var (
		appleSum int
		riceSum float64
		iceScreamSum float64
	)
	
	// type(value) to convert data type for math operation
	appleSum = int(float64(apple) * applePrice)
	fmt.Printf("Cost of apple : %T %v\n", appleSum, appleSum)

	riceSum = (float64(rice) * ricePrice)
	fmt.Printf("Cost of rice : %T %v\n", riceSum, riceSum)

	iceScreamSum = float64(iceScream) * icePrice
	fmt.Printf("Cost of iceScream : %T %v\n", iceScreamSum, iceScreamSum)
}