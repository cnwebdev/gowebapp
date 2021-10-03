// Veriables & Functions
package main

import "fmt"

const birthYear = 1970
const currentYear = 2021

func main() {
	var firstName, lastName string 
		firstName = "Chris"
		lastName = "Nguyen"

	job := "Programmer"

	fmt.Println("Hello", firstName, lastName) 
	fmt.Println(firstName, "is", calcAge(), "years old")
	fmt.Println(firstName, "is a", job)
}

func calcAge() int {
	var age int 
	age = currentYear - birthYear
	return age
}
