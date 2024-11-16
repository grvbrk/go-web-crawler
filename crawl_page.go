package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	urlRaw, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Error parsing rawBaseURL")
		return
	}

	urlCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Error parsing rawCurrentURL")
		return
	}

	if urlCurrent.Hostname() != urlRaw.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Error normalizing current url")
		return
	}

	_, ok := (pages)[normalizedURL]
	if !ok {
		(pages)[normalizedURL] = 1
	} else {
		(pages)[normalizedURL] += 1
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Println("Error getting urls from htmlBody")
		return
	}

	for _, nextURL := range nextURLs {

		crawlPage(rawBaseURL, nextURL, pages)
	}
}
