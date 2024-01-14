package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	str := []string{"One", "Two", "Three", "Four", "Five"}
	Print(ints)
	Print(str)
}
