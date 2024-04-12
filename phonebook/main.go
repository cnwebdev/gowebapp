package main

import (
	"fmt"
	"os"
	"path"
)

type Person struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Person{}

func search(key string) *Person {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s [s|l] <arguments>\n", exe)
		return
	}

	data = append(data, Person{"Mike", "Mcafee", "2021234567"})
	data = append(data, Person{"Time", "Smith", "2025673444"})
	data = append(data, Person{"Kary", "Lee", "2022332567"})

	// Differentiate between the commands
	switch args[1] {
	// the search command
	case "s":
		if len(args) != 3 {
			fmt.Println("Usage: search lastname")
			return
		}
		result := search(args[2])
		if result == nil {
			fmt.Println("Entry not found:", args[2])
			return
		}
		fmt.Println(*result)
	case "l":
		list()
		// Response to anything that is not a match
	default:
		fmt.Println("Not a valid option")
	}
}
