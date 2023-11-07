package main

import "fmt"

var (
	str      = "главрыба"
	reversed = "абырвалг"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Println(str)
	fmt.Println(reverse(str))
	fmt.Println(reverse(str) == reversed)
}
