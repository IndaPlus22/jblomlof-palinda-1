package main

// works on the website but not locally for me again because of lack of imports

import (
	"golang.org/x/tour/wc"
	"strings"
)

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	countOfWords := make(map[string]int)

	for _, word := range words {
		countOfWords[word] += 1
	}

	return countOfWords
}

func main() {
	wc.Test(wordCount)
}
