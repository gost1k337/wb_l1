// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func toSet(w []string) []string {
	res := make([]string, 0)
	tmp := make(map[string]struct{})
	for _, i := range w {
		if _, ok := tmp[i]; !ok {
			tmp[i] = struct{}{}
			res = append(res, i)
		}
	}

	return res
}

func main() {
	w := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(toSet(w))
}
