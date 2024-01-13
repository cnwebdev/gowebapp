// Receivers: Structs with functions
// Types and Structs
package main

import "log"

type Vih struct {
	Make      string
	Model     string
	YearBuilt int
	Type      string
	NumWheel  int
	Color     string
	Weight    int
	NumSeat   int
}

type Person struct {
	FirstName string
	LastName  string
	DoB       string
	Job       string
	Address   string
	Phone     string
	Email     string
}

func main() {

	p1001 := Person{
		FirstName: "Chris",
		LastName:  "Nguyen",
		DoB:       "Augurst, 12, 1970",
		Job:       "Programmer",
		Address:   "1288 Smart St, Sun Rise, FL 33709",
		Phone:     "505-556-9876",
		Email:     "chris@funjob.com",
	}

	maverick := Vih{
		Make:      "Ford",
		Model:     "Maverick 2022",
		YearBuilt: 2021,
		Type:      "Pickup Truck",
		NumWheel:  4,
		Color:     "Area 51",
		NumSeat:   5,
	}

	log.Println("The person first name is:", p1001.printFirstName())

	log.Println(p1001.printFirstName(), "drives the", maverick.printVihMake(), maverick.printVihModel())

}

// function with receiver is a great way to access and return any data field from struct type
func (p *Person) printFirstName() string {
	return p.FirstName
}

func (p *Person) printLastName() string {
	return p.LastName
}

func (p *Person) printDoB() string {
	return p.DoB
}

func (v *Vih) printVihYearBuilt() int {
	return v.YearBuilt
}

func (v *Vih) printVihMake() string {
	return v.Make
}

func (v *Vih) printVihModel() string {
	return v.Model
}
