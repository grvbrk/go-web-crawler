package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s \n", baseURL)
	fmt.Println("=============================")

	type report struct {
		key string
		val int
	}

	pairs := make([]report, 0, len(pages))
	for k, v := range pages {
		pairs = append(pairs, report{key: k, val: v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].val == pairs[j].val {
			return pairs[i].key < pairs[j].key
		}
		return pairs[i].val > pairs[j].val
	})

	for _, p := range pairs {
		fmt.Printf("Found %v internal links to %v\n", p.val, p.key)
	}

}
