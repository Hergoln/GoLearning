package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	histogram := make(map[string]int)
	
	for _, word := range strings.Fields(s) {
		if _, present := histogram[word]; present {
			histogram[word]++
		} else {
			histogram[word] = 1
		}
	}
	
	return histogram
}

func main() {
	wc.Test(WordCount)
}
