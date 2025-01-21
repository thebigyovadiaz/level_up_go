// Challenge - Sleeps until your birthday
// Task: given a target date implement a function that outputs the number
// 		of nights until the given date

package main

import (
	"flag"
	"log"
	"time"
)

var templateFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	pT, err := time.Parse(templateFormat, target)
	if err != nil {
		log.Fatalf("Failed to parse time: %v - target: %s", err, target)
	}
	return pT
}

// calculateSleeps returns the number of sleeps until the target.
func calculateSleeps(target time.Time) float64 {
	untilT := time.Until(target).Hours() / 24
	return untilT
}

func main() {
	birthday := flag.String("birthday", "", "Your next birthday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*birthday)
	sleeps := int(calculateSleeps(target))
	log.Printf("You have %d sleeps until your birthday. Hurray!\n", sleeps)
}
