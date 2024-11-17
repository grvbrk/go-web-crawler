package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	start := time.Now()
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error parsing maxConcurrency cmd argument")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Error parsing maxPages cmd argument")
		os.Exit(1)
	}

	BASE_URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("error parsing baseURL: %v \n", err)
	}

	fmt.Printf("starting crawl of: %s \n", BASE_URL)

	config := config{
		pages:              make(map[string]int),
		baseURL:            BASE_URL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	config.wg.Add(1)
	go config.crawlPage(os.Args[1])
	config.wg.Wait()

	duration := time.Since(start)

	for k, v := range config.pages {
		fmt.Println(k, v)
	}

	fmt.Printf("Time took for crawling %s: %v\n", os.Args[1], duration.Seconds())

	printReport(config.pages, os.Args[1])

}
