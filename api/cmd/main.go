package main

import (
	"goroutine-scraping-api/sites"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		sites.ScrapingHatenaItHotentry()
		defer wg.Done()
	}()

	go func() {
		sites.ScrapingQiitaWeeklyTrend()
		defer wg.Done()
	}()

	wg.Wait()
}
