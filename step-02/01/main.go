// Challenge: Stop Copying Me
// Task: given a list of messages and a number N implement a function
//		that outputs the same message N times concurrently.

package main

import (
	"flag"
	"log"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, msg string) {
	ch := make(chan struct{})

	for i := 0; i < n; i++ {
		go func(i int) {
			log.Printf("[G%d]: %s\n", i, msg)
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < n; i++ {
		<-ch
	}
	close(ch)
}

func main() {
	factor := flag.Int64("factor", 0, "Than fan-out factor to repeat by")
	flag.Parse()
	for _, msg := range messages {
		log.Println(msg)
		repeat(int(*factor), msg)
	}
}
