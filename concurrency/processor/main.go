package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	jobs := make(chan int, 100)
	workerCount := 4

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// spawn normall workers
	for index := 0; index < workerCount; index++ {
		go func(workerNumber int) {
			worker(ctx, workerNumber, jobs)
			wg.Done()
		}(index + 1)
	}

loop:
	for {
		select {
		case <-time.After(time.Millisecond * 100):
			jobs <- rand.Intn(100)
		case <-ctx.Done():
			break loop
		}
	}

	wg.Wait()
	fmt.Println("closing the application")
}

func worker(ctx context.Context, workerId int, jobs <-chan int) {
	for {
		select {
		case job := <-jobs:
			timeoutContext, cancel := context.WithTimeout(ctx, time.Millisecond*250)
			defer cancel()

			processJob(timeoutContext, workerId, job)
		case <-ctx.Done():
			fmt.Printf("worker %d has been closed\n", workerId)
			return
		}
	}
}

func processJob(ctx context.Context, workerId int, job int) {
	duration := time.Duration(200 + rand.Intn(100))

	select {
	case <-time.After(time.Millisecond * duration):
		fmt.Printf("job %d has been processed via worker %d\n", job, workerId)
	case <-ctx.Done():
		fmt.Printf("job %d has been timeout or canceled via worker %d!\n", job, workerId)
	}
}
