// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	c := make(chan int)
	wg := sync.WaitGroup{}

	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c <- i * i
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	sum := 0
	for i := range c {
		sum += i
	}

	fmt.Printf("Sum: %d", sum)

}
