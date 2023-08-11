package sites

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapingQiitaWeeklyTrend() {
	fmt.Println("Qiita週間トレンド")
	url := "https://b.hatena.ne.jp/hotentry/it"

	c := colly.NewCollector()

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		title := e.ChildText("span")
		link := e.ChildAttr("a", "href")
		if strings.HasPrefix(link, "https") {
			fmt.Printf("Qiita週間トレンド タイトル: %s\n", title)
			fmt.Printf("リンク: %s\n\n", link)
		}
	})

	c.Visit(url)
}
