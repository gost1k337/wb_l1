// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersect(s1 []int, s2 []int) []int {
	tmp := make(map[int]struct{})
	res := make([]int, 0)
	for _, i := range s1 {
		if _, ok := tmp[i]; !ok {
			tmp[i] = struct{}{}
			res = append(res, i)
		}
	}

	for _, i := range s2 {
		if _, ok := tmp[i]; !ok {
			tmp[i] = struct{}{}
			res = append(res, i)
		}
	}

	return res
}

func main() {
	s1 := []int{1, 2, 5, 5, 6, 3, 2}
	s2 := []int{2, 3, 4, 4, 1, 9}

	fmt.Println(intersect(s1, s2))
}
