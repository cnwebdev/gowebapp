package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Refactor the previous exercise from the if statement
//  section.
//
//  "Print the number of days in a given month."
//
//  Use a switch statement instead of if statements.
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Enter the month - example [May]")
		return
	}

	// Store the command line args in month variable
	month := args[0]
	month = strings.ToLower(month)

	var (
		jan = "january"   // 31
		feb = "february"  // 28
		mar = "march"     // 31
		apr = "april"     // 30
		may = "may"       // 31
		jun = "june"      // 30
		jul = "july"      // 31
		aug = "august"    // 31
		sep = "september" // 30
		oct = "october"   // 31
		nov = "november"  // 30
		dec = "december"  // 31

		nday int
	)

	year := time.Now().Year()

	fmt.Println("Current year is", year)

	// if statement from previous exercise
	/*
		if (year%4 == 0 || year%400 == 0) && month == feb {
			nday = 29
		} else if month == jan || month == mar || month == may ||
			month == jul || month == aug || month == oct || month == dec {
			nday = 31
		} else if month == apr || month == jun || month == sep || month == nov {
			nday = 30
		} else if month == feb {
			nday = 28
		} else {
			fmt.Printf("%q is not a month\n", month)
			return
		}
	*/

	// Coverted if statement above to switch stetement
	switch m := month; {
	case m == feb:
		if year%4 == 0 || year%400 == 0 {
			nday = 29
		} else {
			nday = 28
		}
	case m == jan, m == mar, m == may, m == jul, m == aug, m == oct, m == dec:
		nday = 31
	case m == apr, m == jun, m == sep, m == nov:
		nday = 30
	default:
		fmt.Printf("%q is not a month\n", month)
		return
	}

	month = strings.Title(month)

	fmt.Printf("%q has %d days\n", month, nday)

}
