// Challenge: Raffle Winner
// Task: given an input.json file with raffle entries implementation a function
//		that outputs a list of structs containing the entries

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const FILE = "./step-01/03/entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	file, err := os.ReadFile(FILE)
	if err != nil {
		log.Fatal(err)
	}

	var data []raffleEntry
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	winner := rand.Intn(len(entries))
	return entries[winner]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is ...")

	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
