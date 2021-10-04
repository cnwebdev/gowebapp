// How to get the length of a utf-8 string
package main

import "fmt"

func main() {
	
	str := "Ho Chu Tich Song Mai Trong Quan Chung Ta"
	// Print the length of str string
	fmt.Println(str, " - length =", len(str))

	str = "Hồ Chủ Tịch Sống Mãi Trong Quần Chúng Ta"
	fmt.Println(str, " - lenght =", len(str))
}