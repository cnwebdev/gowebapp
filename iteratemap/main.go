package main

import "fmt"

func main() {
	aMap := make(map[string]string)
	aMap["January"] = "01"
	aMap["Febuary"] = "02"

	// range works with maps
	for k, v := range aMap {
		fmt.Printf("Key: %s value: %s\n", k, v)
	}
}
