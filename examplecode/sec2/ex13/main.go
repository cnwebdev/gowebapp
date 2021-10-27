// Escape sequences
package main

import "fmt"

func main() {

	// back slash \ is escape sequen using in c and go to allow formating specific character output
	// \n is for inserting a newline
	fmt.Println("Monday\nTuesday\nWednesday\nThusday\nFriday\nSaturday\nSunday")

	// \" will insert double quote
	fmt.Println("The president said \"Make America Great Again!\"")

	// \\ insert backslash in string
	fmt.Println("c:/home/user/name")
	fmt.Println("c:\\home\\user\\name")
}
