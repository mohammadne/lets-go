package basic

import (
	"fmt"
	"time"
)

func simpleGoroutine() {
	work := func(from string) {
		for index := 0; index < 3; index++ {
			fmt.Println(from, ": ", index)
		}
	}

	work("direct")

	go work("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
}
