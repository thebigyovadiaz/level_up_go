// Challenge: Calculate Change
// Task: given an amount and set of coins implement a function that
//		output the optimal solution for the given amount

package main

import (
	"flag"
	"log"
	"math"
)

// Coin contains the name and value of a coin
type Coin struct {
	name  string
	value float64
}

// coins is the list of values available for making change
var coins = []Coin{
	{name: "1 pound", value: 1},
	{name: "50 pence", value: 0.50},
	{name: "20 pence", value: 0.20},
	{name: "10 pence", value: 0.10},
	{name: "5 pence", value: 0.05},
	{name: "1 penny", value: 0.01},
}

// calculateChange return the coins required to calculate the
func calculateChange(amount float64) map[Coin]int {
	change := make(map[Coin]int)

	for _, coin := range coins {
		if amount >= coin.value {
			count := math.Floor(amount / coin.value)
			amount = amount - count*coin.value
			change[coin] = int(count)
		}
	}

	return change
}

// printCoins prints all the coins in the slice to the terminal
func printCoins(change map[Coin]int) {
	if len(change) == 0 {
		log.Println("No change found")
		return
	}

	log.Println("Change has been calculated.")

	for coin, count := range change {
		log.Printf("%d x %s \n", count, coin.name)
	}
}

func main() {
	amount := flag.Float64("amount", 0.0, "The amount you want to make change for")
	flag.Parse()
	change := calculateChange(*amount)
	printCoins(change)
}
