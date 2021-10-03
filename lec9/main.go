// Pointers
package main

import "fmt"

var bYear int = 1970

func main() {

	var cYear = 2021
	
	var fName, lName string
		fName = "Chris"
		lName = "Nguyen"

	job := "Programmer"

	fmt.Println("Hello", fName, lName)

	fmt.Println(fName, "'s age is", cYear - bYear)
	fmt.Println(fName, "'s age after func call is", calcAge(&cYear))
	fmt.Println(fName, "is a", job)

}

// calcAge changes the bYear by pointer
func calcAge(c *int) int {
	*c = 2015
	var age int
	age = *c - bYear
	return age
}