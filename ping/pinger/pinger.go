package pinger

import (
	"context"
	"fmt"
	"log"
	"time"
)

// HandlePing handles the ping event by printing the value to standard output.
func HandlePing(msg string, ping <-chan int, handler func(string, int)) {
	for i := range ping {
		handler(msg, i)
	}
}

// SendPingWithContext sends ping as long as the context has not timed out.
// Once the context is done, close channel and terminate program.
func SendPingWithContext(ctx context.Context, ping chan<- int) {
	deadline, ok := ctx.Deadline()

	if !ok {
		log.Fatal("something went wrong")
	} else {
		fmt.Println("deadline: ", deadline)

		for i := 1; ctx.Err() == nil; i++ {
			ping <- i
			time.Sleep(time.Second)
		}

		fmt.Println("deadline reached: ", time.Now())
	}

}
