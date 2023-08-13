package main

import (
	"goroutine-scraping-cli/sites"

	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	resultChannel := make(chan sites.Article)

	wg.Add(2)

	go func() {
		articles := sites.ScrapingHatenaItHotentry()
		for _, article := range articles {
			resultChannel <- article
		}
		defer wg.Done()
	}()

	go func() {
		articles := sites.ScrapingQiitaWeeklyTrend()
		for _, article := range articles {
			resultChannel <- article
		}
		defer wg.Done()
	}()

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		fmt.Println(result)
	}
}
