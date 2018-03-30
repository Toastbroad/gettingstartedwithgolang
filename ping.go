// Program sends an incrementing integer to the ping channel every second.
// The programs runs a default number of seconds or as many as are set by using the -duration flag.
// Purpose is to get to know the context and flags package, and get hands dirty with channels and go routines.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
)

type pinger struct {
	ping chan int
}

func newPinger() *pinger {
	return &pinger{
		ping: make(chan int),
	}
}

func (pinger *pinger) getPing() <-chan int { // define directionality of returned channel
	return pinger.ping
}

func main() {
	const defaultDuration = 3

	pinger := newPinger()

	// Get duration from -duration flag or use default duration. Flag needs to be parsed, otherwise fall back to defaultDuration.
	duration := flag.Int("duration", defaultDuration, "specify how many seconds to run")
	flag.Parse()

	if *duration <= 0 {
		log.Fatal("duration must be an integer greater than 0")
	}

	//Create context that timeouts after default or specified duration.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*duration))
	defer cancel()

	// Handle the ping event by printing the value to standard output.
	go func() {
		ping := pinger.getPing()
		for i := range ping {
			fmt.Printf("%d \n", i)
		}
	}()

	// As long as the context has not timed out, send ping.
	// Once the context is done, close channel and terminate program.
	for i := 1; ctx.Err() == nil; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("timeout: ", time.Now())
			err := ctx.Err()
			if err != nil {
				fmt.Println("error: ", err)
			}
			close(pinger.ping)
			break
		default:
			pinger.ping <- i
			time.Sleep(time.Second)
		}
	}
}
