// Program sends an incrementing integer to the ping channel every second.
// The programs runs a default number of seconds or as many as are set by using the -duration flag.
// Purpose is to get to know the context and flags package, and get hands dirty with channels and go routines.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"./pinger"
)

func main() {
	const defaultDuration = 3

	ping := make(chan int)

	// Get duration from -duration flag or use default duration. Flag needs to be parsed, otherwise fall back to defaultDuration.
	duration := flag.Int("duration", defaultDuration, "specify how many seconds to run")
	flag.Parse()

	if *duration <= 0 {
		log.Fatal("duration must be an integer greater than 0")
	}

	//Create context that timeouts after default or specified duration.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*duration))
	defer cancel()

	go pinger.HandlePing(ping)

	pinger.SendPing(ctx, ping)
}
