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
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
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

	month = strings.Title(month)

	fmt.Printf("%q has %d days\n", month, nday)

}
