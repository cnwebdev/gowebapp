package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------
const usage = `
Usage: [username] [password]`

func main() {
	var (
		username = "jack"
		password = 1888
	)

	args := os.Args[1:]

	l := len(args)

	if l != 2 || l <= 1 {
		fmt.Println(usage)
		return
	}

	usr, pass := args[0], args[1]

	passwd, err := strconv.Atoi(pass)
	if err != nil {
		fmt.Println("Not a valid entry!")
		fmt.Println(usage)
		return
	}

	if usr != username {
		fmt.Printf("Access denied for %q\n", usr)
	} else if passwd != password {
		fmt.Printf("Invalid password for %q\n", usr)
	} else {
		fmt.Printf("Access granted to %q\n", usr)
	}
}
