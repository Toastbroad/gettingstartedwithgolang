package pinger

import (
	"context"
	"fmt"
	"time"
)

type pinger struct {
	ping chan int
}

func NewPinger() *pinger {
	return &pinger{
		ping: make(chan int),
	}
}

func HandlePing(pinger *pinger) {
	ping := pinger.GetSendPing()
	for i := range ping {
		fmt.Printf("%d \n", i)
	}
}

func SendPing(ctx context.Context, pinger *pinger) {
	for i := 1; ctx.Err() == nil; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("timeout: ", time.Now())
			err := ctx.Err()
			if err != nil {
				fmt.Println("error: ", err)
			}
			close(pinger.GetReceivePing())
			break
		default:
			pinger.GetReceivePing() <- i
			time.Sleep(time.Second)
		}
	}
}

func (pinger *pinger) GetSendPing() <-chan int { // define directionality of returned channel
	return pinger.ping
}

func (pinger *pinger) GetReceivePing() chan<- int { // define directionality of returned channel
	return pinger.ping
}
