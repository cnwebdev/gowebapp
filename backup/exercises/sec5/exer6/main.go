package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// Solve this exercise by using your previous solution for
	// the "Housing Prices" exercise.
	head := strings.Split(header, separator)

	rows := strings.Split(data, "\n")

	var (
		locs                      []string
		size, beds, baths, prices []int
	)

	for _, r := range rows {
		str := strings.Split(r, ",")

		locs = append(locs, str[0])

		n, err := strconv.Atoi(str[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		size = append(size, n)

		n, err = strconv.Atoi(str[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		beds = append(beds, n)

		n, err = strconv.Atoi(str[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		baths = append(baths, n)

		n, err = strconv.Atoi(str[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		prices = append(prices, n)
	}

	// Printing the table header
	for _, v := range head {
		fmt.Printf("%-15s", v)
	}
	fmt.Printf("\n%s\n", strings.Repeat("-", 70))

	// Printing the data table
	for i := range rows {
		fmt.Printf("%-14s %-14d %-14d %-14d %d\n", locs[i], size[i], beds[i], baths[i], prices[i])
	}

	fmt.Printf("%s\n", strings.Repeat("-", 70))

	var ts, tb, tba, tp float64
	for i := range rows {
		ts += float64(size[i])
		tb += float64(beds[i])
		tba += float64(baths[i])
		tp += float64(prices[i])
	}
	l := float64(len(rows))
	fmt.Printf("%-14s %-14.2f %-14.2f %-14.2f %.2f\n", "", ts/l, tb/l, tba/l, tp/l)

}
