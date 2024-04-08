package basic

import (
	"fmt"
	"sync"
	"time"
)

func worker() {
	workerCount := 3
	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	for index := 1; index <= workerCount; index++ {
		go func(index int) {
			wg.Add(1)

			for {
				if job, more := <-jobs; more {
					time.Sleep(time.Millisecond * 200)
					fmt.Printf("job %d procced by worker %d\n", job, index)
				} else {
					wg.Done()
					break
				}
			}
		}(index)
	}

	timer := time.NewTimer(time.Second * 5)
	index := 0

loop:
	for {
		select {
		case <-timer.C:
			close(jobs)
			break loop
		default:
			jobs <- index
			index++
		}
	}

	fmt.Println("All jobs have been finished")
}
