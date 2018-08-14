package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var airportList = map[string]string{
		"BHM": "Birmingham International Airport",
		"DHN": "Dothan Regional Airport",
		"HSV": "Huntsville International Airport",
		"MOB": "Mobile",
		"MGM": "Montgomery",
		"ANC": "Anchorage International Airport",
		"FAI": "Fairbanks International Airport",
		"JNU": "Juneau International Airport",
		"FLG": "Flagstaff",
		"TUS": "Tucson International Airport",
		"YUM": "Yuma International Airport",
		"FYV": "Fayetteville",
		"LIT": "Little Rock National Airport",
		"XNA": "Northwest Arkansas Regional Airport",
		"BUR": "Burbank",
		"FAT": "Fresno",
		"LGB": "Long Beach",
		"LAX": "Los Angeles International Airport",
		"OAK": "Oakland",
		"ONT": "Ontario",
		"PSP": "Palm Springs",
		"SMF": "Sacramento",
		"SAN": "San Diego",
		"SFO": "San Francisco International Airport",
		"SJC": "San Jose",
		"SNA": "Santa Ana",
		"ASE": "Aspen",
		"COS": "Colorado Springs",
		"DEN": "Denver International Airport",
		"GJT": "Grand Junction",
		"PUB": "Pueblo",
		"BDL": "Hartford",
		"HVN": "Tweed New Haven",
		"DCA": "Washington National Airport",
		"DFW": "Dallas, Fort Worth",
	}

	var itinerary []string
	itinerary = askInput()
	var output []string

	for k, v := range itinerary {
		var msg string
		_, legit := airportList[v]
		if legit {
			if k != len(itinerary)-1 {
				msg = airportList[v] + " to "
			} else {
				msg = airportList[v]
			}
		} else {
			msg = "Some unknown airport"
		}
		output = append(output, msg)
	}
	fmt.Println(output)
}

func askInput() []string {
	fmt.Println("Enter airport code, one on each line, Ctrl-D when complete")
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("First Airport code")
	var itinerary []string
	for reader.Scan() {
		fmt.Println("Next Airport Code?: ")
		apc := reader.Text()
		itinerary = append(itinerary, apc)
		if len(apc) < 3 {
			break
		}
		if err := reader.Err(); err != nil {
			break
		}
	}
	return itinerary
}
