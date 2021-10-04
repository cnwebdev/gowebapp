// Path separator example
package main

import (
	"fmt"
	"path"
)


func main() {

	var (
		dir string
		file string
	)

	dir, file = path.Split("home/index.html")
	fmt.Println("Dir and file name are :", dir, file)
	
	// Go allows to discard a return value from function with multiple return values
	_, file = path.Split("home/index.html")
	fmt.Println("Dir and file name are :", file)


}