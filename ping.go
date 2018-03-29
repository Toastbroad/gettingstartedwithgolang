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

func main() {
	const defaultDuration = 3

	ping := make(chan int)

	duration := flag.Int("duration", defaultDuration, "specify how many seconds to run")
	flag.Parse()

	if *duration <= 0 {
		log.Fatal("duration must be an integer greater than 0")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*duration))
	defer cancel()

	go func() {
		for i := range ping {
			fmt.Printf("%d \n", i)
		}
	}()

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
