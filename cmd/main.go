package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("Scraping Start.")
	url := "https://b.hatena.ne.jp/hotentry/it"

	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Title:", e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL.String())
	})

	c.Visit(url)
}
