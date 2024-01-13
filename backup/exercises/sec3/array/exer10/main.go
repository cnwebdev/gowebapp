package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Hipster's Love Bookstore Search Engine
//
//  Your goal is to allow people to search for books.
//
//  1. Create an array with the following book titles:
//      Kafka's Revenge
//      Stay Golden
//      Everythingship
//      Kafka's Revenge 2nd Edition
//
//  2. Get the search query from the command-line argument
//
//  3. Search for the books in the books array
//
//  4. When the program finds the book, print it.
//  5. Otherwise, print that the book doesn't exist.
//
//  6. Handle the errors.
//
// RESTRICTION:
//   + The search should be case insensitive.
//
// EXPECTED OUTPUT
//   go run main.go
//     Tell me a book title
//
//   go run main.go STAY
//     Search Results:
//     + Stay Golden
//
//   go run main.go sTaY
//     Search Results:
//     + Stay Golden
//
//   go run main.go "Kafka's Revenge"
//     Search Results:
//     + Kafka's Revenge
//     + Kafka's Revenge 2nd Edition
//
//   go run main.go void
//     Search Results:
//     We don't have the book: "void"
//
// HINTS:
//   + To find out whether a string contains another string value, you can use the strings.Contains function.
//   + To convert a string value to lowercase, you can use the strings.ToLower function.
//   + Check out the strings package for more information.
// ---------------------------------------------------------

func main() {

	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everthingship",
		"Kafka's Revenge 2nd Edition",
	}

	// Get the search query from the command-line argument
	query := os.Args[1:]

	// Check command-line input syntax
	if len(query) < 1 {
		fmt.Println("Tell me a book title")
		return
	}

	var found bool

	// Search for the books in the books array
	fmt.Println("We found:")
	for _, q := range query {

		// Convert book title into lower case string
		q = strings.ToLower(q)

		for j, b := range books {
			// Convert command-line argument text to lower case
			b = strings.ToLower(b)

			// if the book title found store it in match array
			if strings.Contains(b, q) {

				// Convert the book title back to Title format
				b = strings.Title(b)
				fmt.Printf("#%d: %q\n", j, b)
				found = true
			}
		}
	}

	// Print the search result, other wise print the book doesn't exist.

	if !found {
		fmt.Println(query, "does not exist")
	}
}
