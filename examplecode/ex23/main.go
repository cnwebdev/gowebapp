// Directories Example
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(file string) {
		d := []byte("555AAA555")
		check(os.WriteFile(file, d, 0644))
	}

	createEmptyFile("subdir/file1.txt")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2.txt")
	createEmptyFile("subdir/parent/file3.txt")
	createEmptyFile("subdir/parent/child/file4.txt")

	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
