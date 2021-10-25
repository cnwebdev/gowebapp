// File Paths
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("home", "user", "chris", "text.txt")
	fmt.Println("p:", p)

	a := filepath.Join("home//", "file1.txt")
	b := filepath.Join("home/../home", "doc1.docx")

	fmt.Println(a)
	fmt.Println(b)

	// To find directory or file name user filepath.Dir(path) and filepath.Base(path)
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// To check for absolute directory path = return true or fall
	fmt.Println(filepath.IsAbs("home/user/chris"))
	fmt.Println(filepath.IsAbs("/home/user/chris"))

	// To get file extension use filepath.Ext(filename)
	ext := filepath.Ext(p)
	ext1 := filepath.Ext(a)
	ext2 := filepath.Ext(b)
	fmt.Println(ext, ext1, ext2)

	// To remove file extension use strings.TrimSuffix(filename, ext)
	fmt.Println(strings.TrimSuffix(p, ext))
	fmt.Println(strings.TrimSuffix(a, ext1))
	fmt.Println(strings.TrimSuffix(b, ext2))

	rel, err := filepath.Rel("home/user", "home/user/chris/text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Rel", rel)
}
