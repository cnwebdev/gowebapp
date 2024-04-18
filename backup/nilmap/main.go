package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{}
	fmt.Printf("Printing a map: %v\n", aMap)

	fmt.Println("Assign key/value to the nil map")
	aMap["January"] = 01
	aMap["Febuary"] = 02
	fmt.Printf("Printing map: %v\n", aMap)
	fmt.Printf("Asseccing the map with key 'January': %v\n", aMap["January"])
	fmt.Printf("Asseccing the map with key 'Febuary': %v\n", aMap["Febuary"])
	fmt.Println("Deleting 'Febuary'")
	delete(aMap, "Febuary")
	fmt.Printf("Printing map: %v\n", aMap)

	fmt.Println("Assign aMap to nil")
	aMap = nil
	fmt.Printf("Printing map: %v\n", aMap)
	if aMap == nil {
		fmt.Println("aMap is nil map")
		aMap = map[string]int{}
	}
	fmt.Println("Assign key/value to the map again")
	aMap["January"] = 01
	aMap["January"] = 1
	aMap["Febuary"] = 2
	fmt.Printf("Printing map: %v\n", aMap)
}
