// Challenge: The Big Sale
// Task: given a max budget and a list of sales items implement a function
//		that outputs the sorted list of deals within budget.

package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
)

const PATH = "./step-01/05/items.json"

// SaleItem represents the item part of the big sale
type SaleItem struct {
	Name           string  `json:"name"`
	OriginalPrice  float64 `json:"originalPrice"`
	ReducedPrice   float64 `json:"reducedPrice"`
	SalePercentage float64 `json:"salePercentage"`
}

// matchSales adds the sales percentage of the item
// and sorts the array accordingly
func matchSales(budget float64, items []SaleItem) []SaleItem {
	var matchItem []SaleItem
	for _, item := range items {
		if item.ReducedPrice <= budget {
			item.SalePercentage = -(item.ReducedPrice - item.OriginalPrice) /
				item.OriginalPrice * 100

			matchItem = append(matchItem, item)
		}
	}

	sort.Slice(matchItem, func(i, j int) bool {
		return matchItem[i].SalePercentage > matchItem[j].SalePercentage
	})

	return matchItem
}

func importData() []SaleItem {
	file, err := os.ReadFile(PATH)
	if err != nil {
		log.Fatal(err)
	}

	var data []SaleItem
	if err = json.Unmarshal(file, &data); err != nil {
		log.Fatal(err)
	}

	return data
}

func printItems(items []SaleItem) {
	log.Println("The BIG sale has started with our amazing offers!")

	if len(items) == 0 {
		log.Println("No items found.:( Try increasing your budget.")
	}

	for i, r := range items {
		log.Printf("[%d]:%s is %.2f OFF! Get it now for JUST %.2f!\n",
			i, r.Name, r.SalePercentage, r.ReducedPrice)
	}
}

func main() {
	budget := flag.Float64("budget", 0.0, "The max budget you want to shop with.")
	flag.Parse()

	items := importData()
	matchedItems := matchSales(*budget, items)
	printItems(matchedItems)
}
