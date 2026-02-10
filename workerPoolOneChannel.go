package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Site struct {
	URL string
}

func crawlerOne(workerID int, jobs <-chan Site, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		resp, err := http.Get(job.URL)
		if err != nil {
			fmt.Printf("workerID: %d -- error occurred for URL: %s\n",
				workerID, job.URL)
			continue
		}

		fmt.Printf("workerID: %d -- request success for URL: %s and status code is: %d\n", workerID, job.URL, resp.StatusCode)
		resp.Body.Close()
	}

}

func workerPoolOneChannel() {

	urls := []string{
		"https://www.google.com",
		"https://duckduckgo.com/",
		"https://www.bing.com/",
		"https://www.startpage.com/",
		"https://www.ecosia.org/",
		"https://www.qwant.com/",
		"https://www.kagi.com/",
	}

	jobsChannel := make(chan Site)
	var wg sync.WaitGroup

	// initialte crawlers
	for i := range 3 {
		wg.Add(1)
		go crawlerOne(i+1, jobsChannel, &wg)
	}

	// send jobs
	for _, url := range urls {
		jobsChannel <- Site{URL: url}
	}
	close(jobsChannel)

	wg.Wait()
}
