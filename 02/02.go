// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}

	wg.Wait()
}
