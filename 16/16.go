package main

import (
	"fmt"
	"math/rand"
)

// 16, 4, 3, 0, 7, 1, -2, 10, 8, -5

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// выбираем опорный элемент
	pivotIndex := rand.Int() % len(a)

	// меняем местами опорный элемент с правым
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// отбрасываем меньшие опорного левее
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// ставим опорный сразу после последнего меньшего
	a[left], a[right] = a[right], a[left]

	// применяем qsort к левой и правой частям от опорного
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}

func main() {
	arr := []int{16, 4, -5, 0, 7, 1, -2, 10, 8, 3}
	fmt.Println(arr)
	qsort(arr)
	fmt.Println(arr)
}
