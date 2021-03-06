package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	// only 1 thread in this application
	// but with asynchronously processing the logic
	ch := make(chan string, len(words))

	// Messages themselves could be put in a kind of limbo.
	// The buffer allows them to wait in the channel
	// until a receiver was ready to process them
	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}
