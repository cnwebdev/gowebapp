// Types and Structs
package main

import "log"

type Vih struct {
	Make      string
	Model     string
	YearBuilt string
	Type      string
	NumWheel  string
	Color     string
	Weight    string
	NumSeat   string
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
		YearBuilt: "2021",
		Type:      "Pickup Truck",
		NumWheel:  "4",
		Color:     "Area 51",
		NumSeat:   "5",
	}

	log.Println("Engineer name:", getInfo(p1001.FirstName), getInfo(p1001.LastName))
	log.Println("Date of Birth:", getInfo(p1001.DoB))
	log.Println("Job:", getInfo(p1001.Job))

	log.Println("What vihicle does this person drive?", getInfo(maverick.Make), "Model:", getInfo(maverick.Model))
	log.Println("The vihicle year build:", getInfo(maverick.YearBuilt))

}

func getInfo(str string) string {
	return str
}
