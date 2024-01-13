package main

import (
	"fmt"
	"time"
)

func main() {

	year := time.Now().Year()
	month := time.Now().Month()

	time1 := time.Now().Hour()

	fmt.Println(year, month, time1)
}
