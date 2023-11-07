package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {

	uniqueSymbols := make(map[rune]struct{})

	for _, v := range strings.ToLower(str) {

		if _, ok := uniqueSymbols[v]; ok {
			return false

		} else {
			uniqueSymbols[v] = struct{}{}
		}
	}

	return true
}

func main() {
	str := "abc"

	fmt.Println(checkUnique(str))
}
