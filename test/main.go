// Testing the for loop
package main

import "fmt"

func main() {

	var (
		i   = 1
		sum int
	)

	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}
		sum += i

		fmt.Println(i, "-->", sum)
		i++
	}

}
