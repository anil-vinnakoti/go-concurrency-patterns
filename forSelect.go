package main

import (
	"fmt"
	"time"
	"context"
)

func worker(ctx context.Context, channel <-chan int){
	for{
		select{
		case job, ok := <- channel:
			if !ok{
				fmt.Println("jobs channel closed, worked exiting")
				return
			}
			fmt.Println("processing job", job)
			time.Sleep(500 * time.Millisecond)

		case <-ctx.Done():
			fmt.Println("context cancelled, worker exiting")
			return
		}
	}
}

func forSelect(){
	jobsChannel := make(chan int)
	ctx, contextCancel := context.WithCancel(context.Background())

	go worker(ctx, jobsChannel)

	for i:=1; i<=5; i++{
		jobsChannel <- i
	}

	time.Sleep(time.Second)
	contextCancel()
	close(jobsChannel)
}