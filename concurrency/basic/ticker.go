package basic

import (
	"fmt"
	"time"
)

func timeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

func cronJob() {
	ticker := time.NewTicker(time.Millisecond * 100)
	done := make(chan bool)

	go func() {
		for {
			select {
			case tick := <-ticker.C:
				fmt.Println("Tick at", tick)
			case <-done:
				return
			}
		}
	}()

	time.Sleep(time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
