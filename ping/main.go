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

	"./pinger"
)

const defaultDuration = 3

var duration = flag.Int("duration", defaultDuration, "specify how many seconds to run")

func main() {

	flag.Parse()

	ping := make(chan int)

	if *duration <= 0 {
		log.Fatal("duration must be an integer greater than 0")
	}

	//Create context that timeouts after default or specified duration.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*duration))
	defer cancel()

	go pinger.HandlePing("pong #1", ping, func(msg string, i int) { fmt.Printf("%s %d \n", msg, i) })
	go pinger.HandlePing("pong #2", ping, func(msg string, i int) { fmt.Printf("%s %d \n", msg, i) })

	pinger.SendPingWithContext(ctx, ping)
}
