// Veriable Declaration Example
package main

import "fmt"

// City with highest population
var HiPopCity string

// Team with highest score 
var ChampTeam string

func main() {

	// Grouping related variable
	var (
		topTenCity string
		sortResult string
	)

	/* Exmaple population data
	1. tokyo = 37435191
	2. delhi = 29399141
	3. shanghai = 26317104
	4. saoPaulo = 21846507
	5. mexicoCity = 21671908
	6. cairo = 20484965
	7. dhaka = 20383552
	8. mumbai = 20185064
	9. beijing = 20035445 
	10. Osaka = 19222665
	*/

	topTenCity = "Tokyo, Delhi, Shanghai, Sao Paulo, Mexico City, Cairo, Dhaka, Mumbai, Beijing, Osaka" 
	sortResult = "Tokyo, Delhi, Shanghai, Sao Paulo, Mexico City, Cairo, Dhaka, Mumbai, Beijing, Osaka" 

	// top college footbal rankins AP poll
	var (
		topTenTeam string
		sortList string
	)
	
	/* Example team score
	1. alabama = 1541
	2. georgia = 1497
	3. iowa = 1381
	4. pennState = 1360
	5. cincinnati = 1320
	6. oklahoma = 1248
	7. ohioState = 1094
	8. oregon = 1069
	9. Michigan = 1053
	10. BYU = 990
	*/
	topTenTeam = "Alabama, Georgia, Iowa, PennState, Cincinnati, Oklahoma, Ohio State, Oregon, Michigon, BYU"
	sortList = "Alabama, Georgia, Iowa, PennState, Cincinnati, Oklahoma, Ohio State, Oregon, Michigon, BYU"

	HiPopCity = "Tokyo"
	ChampTeam = "Alabama"
	fmt.Println("Top ten city with by population are", topTenCity, sortResult)
	fmt.Println("Top ten team with most scrore are", topTenTeam, sortList)
	fmt.Println("The City with highest poplulation is", HiPopCity)
	fmt.Println("The Team with top score is", ChampTeam)
}