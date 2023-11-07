// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	var n uint
	fmt.Printf("Введите время действия программы: ")
	fmt.Scanln(&n)
	duration := time.Duration(n) * time.Second

	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	go func() {
		for {
			fmt.Printf("Received: %d\n", <-ch)
		}
	}()

	time.Sleep(duration)
}
