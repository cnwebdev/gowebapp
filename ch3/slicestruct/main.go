package main

import "fmt"

type Record struct {
	FirstName string
	LastName  string
	Phone     string
}

func main() {
	s := []Record{}
	thanh := Record{FirstName: "Thanh", LastName: "Nguyen", Phone: "209-345-3456"}
	xuan := Record{FirstName: "Xuan", LastName: "Nguyen", Phone: "209-256-9887"}
	s = append(s, thanh, xuan)

	// Accessing the fields of the first element
	fmt.Println("index 0:", s[0].FirstName, s[1].LastName, s[0].Phone)
	fmt.Println("Number of structures:", len(s))

	for _, v := range s {
		fmt.Println("First Name:", v.FirstName, "Last Name:", v.LastName, "Phone:", v.Phone)
	}
}
