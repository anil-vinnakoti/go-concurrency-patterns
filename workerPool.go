package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerPool() {
	userIDs := make([]int, 100)
	for i := range userIDs {
		userIDs[i] += i + 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	numsOfWorkers := 10
	jobChannel := make(chan int)

	// start workers
	for i := 0; i < numsOfWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("printing before context cancelling...")
					return
				case jobID, ok := <-jobChannel:
					if !ok {
						return
					}
					processUser(ctx, jobID, workerID)
				}
			}

		}(i)

	}

	// send jobs
	for _, userID := range userIDs {
		jobChannel <- userID
	}

	close(jobChannel) // important
	wg.Wait()
	fmt.Println("printing after waiting from waitGroup...")
}

func processUser(ctx context.Context, jobID int, workerID int) {
	fmt.Printf("worker %d processing user %d\n", workerID, jobID)
	time.Sleep(time.Millisecond * 200)
}
