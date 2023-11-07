package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SyncCounter struct {
	value int64
}

func (s *SyncCounter) IncValue(amount int64) {
	atomic.AddInt64(&s.value, amount)
}

func (s *SyncCounter) GetValue() int64 {
	return atomic.LoadInt64(&s.value)
}

func main() {
	counter := SyncCounter{}

	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.IncValue(1)
		}()
	}

	wg.Wait()

	fmt.Println(counter.GetValue())
}
