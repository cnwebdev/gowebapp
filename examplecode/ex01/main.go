// Introduction to Veriables

package main

import "fmt"

// Exportable variable start with Capital leter
var CamaroSpeed = 195

func main() {

	// Go support camelcase naming convention
	var usPop int
	usPop = 333431917 

	fmt.Println("US population as of Oct 3, 2021 is", usPop)

	fmt.Println("The camaro stop speed is", CamaroSpeed, "Miles per hour")
}

// Go naming rules

/* 
These are not allow in Go

var 3speed int // veriable name start with a number

var !speed int // start with a special character

var spe!ed int // contain special charater

var var int // has same name with Go reserved keyword 

 */