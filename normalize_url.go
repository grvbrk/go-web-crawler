package main

import (
	"fmt"
	"net/url"
	"strings"
)

// func normalizeURL(url string) (string, error) {
// 	ok := strings.HasPrefix(url, "https://")
// 	if ok {
// 		str, ok := strings.CutPrefix(url, "https://")
// 		if !ok {
// 			return "", fmt.Errorf("error while cutting the prefix")
// 		}
// 		hasTrailingSlash := strings.HasSuffix(str, "/")
// 		if hasTrailingSlash {
// 			str, ok := strings.CutSuffix(str, "/")
// 			if !ok {
// 				return "", fmt.Errorf("error while cutting the suffix")
// 			}
// 			return str, nil
// 		}
// 		return str, nil
// 	}

// 	ok = strings.HasPrefix(url, "http://")
// 	if ok {
// 		str, ok := strings.CutPrefix(url, "http://")
// 		if !ok {
// 			return "", fmt.Errorf("error while cutting the prefix")
// 		}
// 		hasTrailingSlash := strings.HasSuffix(str, "/")
// 		if hasTrailingSlash {
// 			str, ok := strings.CutSuffix(str, "/")
// 			if !ok {
// 				return "", fmt.Errorf("error while cutting the suffix")
// 			}
// 			return str, nil
// 		}
// 		return str, nil
// 	}
// 	return "", fmt.Errorf("invalid url")
// }

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid url")
	}
	path, ok := strings.CutSuffix(u.Path, "/")
	if !ok {
		return u.Host + u.Path, nil
	}

	return u.Host + path, nil

}
