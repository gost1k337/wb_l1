package main

import "fmt"

func fillIn(s []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, i := range s {
			out <- i
		}
	}()

	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			out <- i * 2
		}
	}()

	return out
}

func main() {
	s := []int{2, 4, 5, 7, 10, 11, 13}
	for i := range double(fillIn(s)) {
		fmt.Printf("%d\n", i)
	}
}
