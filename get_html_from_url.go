package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	document, err := html.Parse(strings.NewReader(htmlBody))
	results := []string{}
	if err != nil {
		return nil, fmt.Errorf("error parsing out htmlBody")
	}
	for node := range document.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, a := range node.Attr {
				if a.Key == "href" {
					urlPath := a.Val
					ok := strings.HasPrefix(urlPath, "https")
					if !ok {
						results = append(results, rawBaseURL+urlPath)
					} else {
						results = append(results, urlPath)
					}
					break
				}
			}
		}
	}
	return results, nil
}
