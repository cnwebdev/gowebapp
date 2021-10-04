// How to get the length of a utf-8 string
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	
	str := "Ho Chu Tich Song Mai Trong Quan Chung Ta"
	// Print the number of bytes in english ascii code of str string
	fmt.Println(str, " - bytes count =", len(str))

	// To count the number of bytes in utf-8 code from the string 
	str = "Hồ Chủ Tịch Sống Mãi Trong Quần Chúng Ta"
	fmt.Println(str, " - bytes count =", len(str))

	// To count the number of rune in the utf-8 string
	fmt.Println(str, " - rune count", utf8.RuneCountInString(str))
	
}