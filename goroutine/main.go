package main

import (
	"fmt"
	"time"
)

// Function to reverse a word and display each step with a 3-second interval
func reverseWord(word string) {
	n := len(word)
	reversed := make([]byte, n)

	for i := 0; i < n; i++ {
		reversed[i] = word[n-1-i]
		fmt.Println(string(reversed[:i+1]))
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go reverseWord("phone")

	// Allow time for the goroutine to finish
	time.Sleep(18 * time.Second)
}
