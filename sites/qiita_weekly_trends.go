package sites

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapingQiitaWeeklyTrend() []Article {
	fmt.Println("Qiita週間トレンド")
	url := "https://b.hatena.ne.jp/hotentry/it"

	c := colly.NewCollector()

	var articles []Article

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		title := e.ChildText("span")
		link := e.ChildAttr("a", "href")
		if strings.HasPrefix(link, "https") {
			article := Article{Title: title, URL: link}
			articles = append(articles, article)
		}
	})

	c.Visit(url)

	return articles
}
