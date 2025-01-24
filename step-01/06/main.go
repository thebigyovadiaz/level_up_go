// Challenge: Giggest Market
// Task: given a list, users.json and their details implement a function
//		that outputs the country with the most users.json.

package main

import (
	"encoding/json"
	"log"
	"os"
)

const FILE = "./step-01/06/users.json"

type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func getBiggestMarket(users []User) (string, int) {
	counts := make(map[string]int)

	for _, user := range users {
		counts[user.Country]++
	}

	maxCountry := ""
	maxCount := 0

	for country, count := range counts {
		if count > maxCount {
			maxCountry = country
			maxCount = count
		}
	}

	return maxCountry, maxCount
}

func importData() []User {
	file, err := os.ReadFile(FILE)
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	if err := json.Unmarshal(file, &users); err != nil {
		log.Fatal(err)
	}

	return users
}

func main() {
	usersData := importData()
	country, count := getBiggestMarket(usersData)
	log.Printf("The biggest user market is: %s with %d users.", country, count)
}
