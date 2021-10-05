// the different between Prinln vs Printf
package main

import "fmt"

func main() {
	
	var brand string
	// Printf allows printing output with verb parameters like %q and many other
	// q is priting "" around the data
	fmt.Printf("%q\n", brand) // output will be ""

	// assign a string to brand 
	brand = "Golang is cool!"

	// Now print with %q again
	fmt.Printf("%q\n", brand) // => "Golang is cool!"

	// -----------------------------------------------
	// Now take a look at Println

	var (
		tundra = 39999
		ownder = "Chris Ng."
		engStart = true
	)

	// Println example
	fmt.Println("Toyota Tundra 2021 price is", tundra)
	fmt.Println("The owner is", ownder)
	fmt.Println("The Tundra 2021 cost:", tundra, "/ Owner:", ownder, "/ Engine Start:", engStart)

	// Here is the Printf example
	fmt.Printf("The Tundra 2021 cost: %v / Owner: %v / Engine Start: %v\n", tundra, ownder, engStart)

}