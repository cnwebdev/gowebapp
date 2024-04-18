package main

import "fmt"

type Person struct {
	Name    string
	SurName string
	BYear   int
	Age     int
}

// Initialized by return a copy of the object
func NewPerson() Person {
	return Person{}
}

// Initialized by returning a pointer of the object
func New() *Person {
	return &Person{}
}

func main() {
	person1 := NewPerson() // copy of Person
	fmt.Printf("Default value for person1 is %v\n", person1)
	person1.Name = "Thanh"
	person1.SurName = "Nguyen"
	fmt.Printf("Name and surname assigned, %v\n", person1)
	person2 := New() // pointer to Person
	fmt.Printf("Default value for person2 is %v\n", person2)
	person2.Name = "Kim"
	person2.SurName = "Nguyen"
	fmt.Printf("Name and surname assigned, %v\n", person2)
}
