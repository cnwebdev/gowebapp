package main

import "fmt"

func main() {
	// Working with runes
	a := "Đồng Nai Việt Nam"
	fmt.Println(a)

	// Printing the a string as rune
	fmt.Println("Print the above string to runes")
	for _, r := range a {
		fmt.Printf("%v ", r)
	}
	fmt.Println()

	b := 'Đ'
	// Print out rune value
	fmt.Printf("Rune of Đ is %v\n", b)

	// Converting rune to string
	fmt.Printf("Convert rune %v to string %s\n", b, string(b))
}
