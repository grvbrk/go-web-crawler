package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}

	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if len(cfg.pages) >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Error normalizing current url")
		return
	}

	_, ok := cfg.pages[normalizedURL]
	if !ok {
		cfg.mu.Lock()
		cfg.pages[normalizedURL] = 1
		cfg.mu.Unlock()
	} else {
		cfg.mu.Lock()
		cfg.pages[normalizedURL] += 1
		cfg.mu.Unlock()
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Println("Error getting urls from htmlBody")
		return
	}

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
