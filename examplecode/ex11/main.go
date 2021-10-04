// IOTA
package main

import "fmt"

func main() {
	const (
		monday = 1 + iota
		tuesday 
		wednesday
		thursday
		friday 
		saturday 
		sunday 
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}