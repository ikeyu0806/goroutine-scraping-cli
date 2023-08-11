package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("はてなITホットエントリー")
	url := "https://b.hatena.ne.jp/hotentry/it"

	c := colly.NewCollector()

	c.OnHTML(".entrylist-contents-main", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		link := e.ChildAttr("a", "href")
		fmt.Printf("タイトル: %s\n", title)
		fmt.Printf("リンク: %s\n\n", link)
	})

	c.Visit(url)
}
