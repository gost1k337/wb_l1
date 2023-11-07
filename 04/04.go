// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
// произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WorkerRun(id int, ch chan int) {
	go func() {
		for work := range ch {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(fmt.Sprintf("Worker %d: %d", id, work))
		}
	}()
}

func main() {
	var n int
	ch := make(chan int)

	fmt.Printf("Введите количество воркеров: ")
	fmt.Scanln(&n)

	// запись сигналов ос
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	for i := 0; i < n; i++ {
		WorkerRun(i, ch)
	}

loop:
	for {
		select {
		// ждем сигнал от ос и завершаем программу
		case <-sig:
			{
				close(ch)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Shutting down...")
				break loop
			}
		default:
			ch <- rand.Int()
		}

	}
}
