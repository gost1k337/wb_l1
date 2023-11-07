package main

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	sync.RWMutex
	data map[int]int
}

func (sm *SyncMap) Set(k int, v int) {
	sm.Lock()
	sm.data[k] = v
	sm.Unlock()
}

func (sm *SyncMap) Get(k int) (int, bool) {
	sm.RLock()
	v, ok := sm.data[k]
	sm.RUnlock()
	return v, ok
}

func main() {
	sm := SyncMap{data: make(map[int]int)}

	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(i%10, i)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v, _ := sm.Get(i % 10)
			fmt.Printf("%d: %d\n", i, v)
		}(i)
	}
	wg.Wait()
}
