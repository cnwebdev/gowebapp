package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("please provide an argument!")
		return
	}

	file := args[1]
	// Get the OS active environment path
	path := os.Getenv("PATH")
	// split the path into a string slice of directories
	pathSplit := filepath.SplitList(path)
	for _, d := range pathSplit {
		// Join the directory path and the executable file name
		fullPath := filepath.Join(d, file)
		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			// get the file mode
			mode := fileInfo.Mode()
			// Is it a regular file?
			if mode&0111 != 0 {
				fmt.Println(mode, fullPath)
			}
		}
	}
}
