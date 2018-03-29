package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	ping := make(chan int)

	duration := flag.Int("duration", 3, "specify how many seconds to run")
	flag.Parse()

	if *duration <= 0 {
		log.Fatal("WTF")
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
