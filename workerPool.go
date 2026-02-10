package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type WebSite struct {
	URL string
}

type Result struct {
	workerID int
	URL      string
	Status   int
}

func crawler(workedID int, jobs <-chan WebSite, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for job := range jobs {
		resp, err := client.Get(job.URL)
		if err != nil {
			result <- Result{URL: job.URL, Status: 400, workerID: workedID}
			continue
		}
		result <- Result{URL: job.URL, Status: resp.StatusCode, workerID: workedID}
		resp.Body.Close()
	}
}

func workerPool() {
	var wg sync.WaitGroup

	jobsChannel := make(chan WebSite)
	resultChannel := make(chan Result)

	urls := []string{
		"https://www.google.com",
		"https://duckduckgo.com/",
		"https://www.bing.com/",
		"https://www.startpage.com/",
		"https://www.ecosia.org/",
		"https://www.qwant.com/",
		"https://www.kagi.com/",
	}

	// start workers
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go crawler(i, jobsChannel, resultChannel, &wg)
	}

	// close result channel after workers finish
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// send jobs
	go func() {
		for _, url := range urls {
			jobsChannel <- WebSite{URL: url}
		}
		close(jobsChannel)
	}()

	// collect results
	for result := range resultChannel {
		fmt.Printf("Worker %d processed %s with status %d\n",
			result.workerID, result.URL, result.Status)
	}

}
