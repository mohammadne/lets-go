package basic

import (
	"sync"
	"sync/atomic"
)

func incementWithRace(additions, goroutines int) uint64 {
	var value uint64
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < additions/goroutines; c++ {
				value++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return value
}

func incementWithoutRace1(additions, goroutines int) uint64 {
	var value atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < additions/goroutines; c++ {
				value.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return value.Load()
}

func incementWithoutRace2(additions, goroutines int) uint64 {
	var value uint64
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < additions/goroutines; c++ {
				mutex.Lock()
				value++
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return value
}
