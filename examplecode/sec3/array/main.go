// Array
package main

import "fmt"

func main() {

	const max = 100

	var arr = [max]int{}
	var sum int

	for i := 0; i <= max-1; i++ {
		arr[i] = i + 1
	}

	for i, a := range arr {
		fmt.Printf("#%-5d %-5d\n", i, a)
	}
	fmt.Println()

	for i, n := range arr {
		sum += n
		fmt.Printf("#%-5d %d\n", i, sum)
	}
}
