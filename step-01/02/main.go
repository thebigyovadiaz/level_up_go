// Challenge: Slow Down
// Task: implement a function that takes in a string and outputs
//		a prolonged version of each word

package main

import (
	"log"
	"strings"
	"time"
)

const DELAY = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(DELAY)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	for _, word := range words {
		var parseWord []string
		for i, c := range word {
			repeatLetter := strings.Repeat(string(c), i+1)
			parseWord = append(parseWord, repeatLetter)
		}

		print(strings.Join(parseWord, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
