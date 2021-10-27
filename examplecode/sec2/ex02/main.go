// Data Type - Go support static data type

package main

import "fmt"


func main() {

	// declare a veriable with type float64, the value of ricePrice is 0 
	// Go assign default value of 0(zero) for int, uint, uint64, float, float64 
	var ricePrice float64 
	var gasPrice float32
	var carSpeed int
	var airPlanSpeed int64
	var width uint
	var height uint64
	
	fmt.Printf("rice price     : %T %v\n", ricePrice, ricePrice)
	fmt.Printf("gas price      : %T %v\n", gasPrice, gasPrice)
	fmt.Printf("car speed      : %T %v\n", carSpeed, carSpeed)
	fmt.Printf("air Plan Speed : %T %v\n", airPlanSpeed, airPlanSpeed,)
	fmt.Printf("with           : %T %v\n", width, width)
	fmt.Printf("height         : %T %v\n", height, height)

	// Default value for string = " " empty string
	var firstName string
	fmt.Printf("Devault value for veriable of type string is : %T %v\n", firstName, firstName)

	// Default value for boolean is false 
	var switchOn bool
	fmt.Printf("Boolean default value is %T %v\n", switchOn, switchOn)
	switchOn = true
	fmt.Printf("Boolean value after assignment is %T %v\n", switchOn, switchOn)

	// Go does not allow using variable before variable declaration
	// This will not work fmt.Println(lastName)
	lastName := "Chris"
	
	fmt.Println(lastName) // this should work 


}