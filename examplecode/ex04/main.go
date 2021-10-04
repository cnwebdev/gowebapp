// Short declaration in Go
package main

import "fmt"

// package scope variable example
var gameScore int
var searchResult string 

// Don't use short declaration for package scope
// CurYear := 2021 // this will not compile

func main() {
	
	// Where to use short declaration?
	// Here is a convention - if you don't know the initial value, use normal declaration, or 
	// if you know the initial value, then use short declaration

	var product int // veriable to hold result of num1 * num2

	num1 := 2009
	num2 := 389

	product = num1 * num2

	fmt.Println("The product of 2009 and 389 is", product)

	// fmt.Println("Current year is", CurYear)
}