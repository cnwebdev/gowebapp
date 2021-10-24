package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func main() {

	str := os.Args[1:]

	// fmt.Println(len(str), str)

	l := len(str)

	if l < 1 || len(str[0]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	s := str[0]

	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" || s == "y" {
		fmt.Printf("%q is %s\n", s, "vowel")
	} else {
		fmt.Printf("%q is %s\n", s, "consolant")
	}

}
