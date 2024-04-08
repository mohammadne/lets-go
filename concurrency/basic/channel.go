package basic

import (
	"fmt"
	"time"
)

func simpleChannel() {
	pipeline := make(chan struct{ Text string }, 2)

	go func() {
		pipeline <- struct{ Text string }{"msg1 from goroutine"}
		pipeline <- struct{ Text string }{"msg2 from goroutine"}
	}()

	msg1 := <-pipeline
	fmt.Println("one message has been arrived")
	msg2 := <-pipeline

	fmt.Println(msg1.Text + " and " + msg2.Text)
}

func synchChannel() {
	work := func(done chan<- struct{}) {
		time.Sleep(time.Second)
		done <- struct{}{}
	}

	done := make(chan struct{})
	go work(done)

	<-done
	fmt.Println("done")
}

func nonBlocking() {
	messages := make(chan string)

	// non blocking recieve
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// non blocking send
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
}

func simulatePingPong() {
	pingPipeline, pongPipeline := make(chan int), make(chan int)
	done := make(chan struct{})

	go func() {
		for pong := range pongPipeline {
			fmt.Printf("Pong %d recieved\n", pong)
		}
		done <- struct{}{}
	}()

	go func() {
		for {
			if ping, more := <-pingPipeline; more {
				fmt.Printf("Ping %d recieved\n", ping)
				pongPipeline <- ping
			} else {
				close(pongPipeline) // no more value will be send to this channel
				return
			}
		}
	}()

	for index := 0; index < 5; index++ {
		pingPipeline <- index
	}
	close(pingPipeline) // no more value will be send to this channel

	<-done
	fmt.Println("simulation has been done")
}
