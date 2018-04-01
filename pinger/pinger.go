package pinger

import (
	"context"
	"fmt"
	"time"
)

// HandlePing handles the ping event by printing the value to standard output.
func HandlePing(ping <-chan int) {
	for i := range ping {
		fmt.Printf("%d \n", i)
	}
}

// SendPing sends ping as long as the context has not timed out.
// Once the context is done, close channel and terminate program.
func SendPing(ctx context.Context, ping chan<- int) {
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
