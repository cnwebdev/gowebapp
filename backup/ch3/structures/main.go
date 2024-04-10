package main

import "fmt"

// Knowing the data structures of a program is really important
// Programs = Data Structures + Algorithms
type Entry struct {
	Name     string
	LastName string
	Year     int
}

// Initialized by Go
func zeroS() Entry {
	return Entry{}
}

// Initialized by the user
func initS(N, L string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, LastName: L, Year: 2000}
	}
	return Entry{Name: N, LastName: L, Year: Y}
}

// Initialized by Go - returns pointer
func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}

// Initialized by the user - return pointer
func initPtoS(N, L string, Y int) *Entry {
	if len(L) == 0 {
		return &Entry{Name: N, LastName: "Unknown", Year: Y}
	}
	return &Entry{Name: N, LastName: L, Year: Y}
}

func main() {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "p1:", *p1)

	s2 := initS("Thanh", "Nguyen", 2020)
	p2 := initPtoS("Thanh", "Nguyen", 2020)
	fmt.Println("s2:", s2, "p2", *p2)
	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)
	pS := new(Entry)
	fmt.Println("pS:", pS)
}
