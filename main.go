package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	BASE_URL := os.Args[1]
	fmt.Printf("starting crawl of: %s", BASE_URL)
	html, err := getHTML(BASE_URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(html)
}

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("status code %v", res.StatusCode)
	}

	if !strings.HasPrefix(res.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("content-type is not text/html")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil

}
