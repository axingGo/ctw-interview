package scraper

import (
	"ctx-interview/pkg/storage"
	"fmt"

	"github.com/gocolly/colly"
)

type Task struct {
	Name    string            `json:"name"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type Scraper struct{}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) Scrape(task Task) ([]*storage.HotelInfo, error) {
	var results []*storage.HotelInfo
	c := colly.NewCollector()

	// 设置回调函数，提取房源名称和房客推荐
	c.OnHTML("div.lxq01kf", func(e *colly.HTMLElement) {
		// 提取房源名称
		propertyName := e.ChildText("span.a8jt5op") // 这里通过指定的类名获取房源名称
		if propertyName != "" {
			fmt.Println("房源名称: ", propertyName)
		}
	})

	err := c.Visit(task.URL)
	return results, err
}
