package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("No enough arguments")
		return
	}

	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range args[1:] {
		// Is it an integer?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		// Is it a float
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		// Then it is invalid
		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)
	if len(invalid) > total {
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
