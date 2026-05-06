package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	count := make(map[string]int)
	
	for _, word := range fields {
		curr, ok := count[word]
		
		if !ok {
		  count[word] = 1
		} else {
		  count[word] = curr + 1
		}
	}
	
	return count
}

func main() {
	wc.Test(WordCount)
}

