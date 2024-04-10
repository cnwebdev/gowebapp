// Get user input with os.Args
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: <command> <user input>")
		os.Exit(0)
	}

	var min, max float64
	for i, v := range args[1:] {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			_ = fmt.Errorf("failed to convert float %v", err)
		}
		if i == 0 {
			min, max = n, n
			continue
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}
	fmt.Printf("Min %.2f, Max %.2f\n", min, max)
}
