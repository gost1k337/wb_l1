package main

import (
	"fmt"
	"strings"
)

var (
	words    = "snow dog sun"
	reversed = "sun dog snow"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(words)
	fmt.Println(reverseWords(words))
	fmt.Println(reversed == reverseWords(words))
}
