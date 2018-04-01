package pinger

import (
	"context"
	"fmt"
	"time"
)

// HandlePing handles the ping event by printing the value to standard output.
func HandlePing(msg string, ping <-chan int, handler func(string, int)) {
	for i := range ping {
		//fmt.Printf("%s %d \n", msg, i)
		handler(msg, i)
	}
}

// SendPingWithContext sends ping as long as the context has not timed out.
// Once the context is done, close channel and terminate program.
func SendPingWithContext(ctx context.Context, ping chan<- int) {
	deadline, ok := ctx.Deadline()
	fmt.Println("Deadline: ", deadline)
	fmt.Println("Ok: ", ok)
	for i := 1; ctx.Err() == nil; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("timeout: ", time.Now())
			err := ctx.Err()
			if err != nil {
				fmt.Println("error: ", err)
			}
			close(ping)
			break
		default:
			ping <- i
			time.Sleep(time.Second)
		}
	}
}
