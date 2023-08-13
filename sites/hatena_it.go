package sites

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func ScrapingHatenaItHotentry() []Article {
	fmt.Println("はてなITホットエントリー")
	url := "https://b.hatena.ne.jp/hotentry/it"

	c := colly.NewCollector()

	var articles []Article

	c.OnHTML(".entrylist-contents-main", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		link := e.ChildAttr("a", "href")
		article := Article{Title: title, URL: link}
		articles = append(articles, article)
	})

	c.Visit(url)

	return articles
}
