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
	c := colly.NewCollector()

	// 存放 href 值的切片
	var hrefs []string

	// 查找 a 标签，rel 等于 "noopener noreferrer nofollow"
	c.OnHTML(`a[rel="noopener noreferrer nofollow"]`, func(e *colly.HTMLElement) {
		// 提取 href 属性
		href := e.Attr("href")
		if href != "" {
			hrefs = append(hrefs, href)
		}
	})

	// 处理请求错误
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request failed:", err)
	})
	err := c.Visit(task.URL)

	// 到详情页获取hotel具体信息
	var hotelInfos []*storage.HotelInfo
	// for _, href := range hrefs {
	// 	// todo
	// }
	return hotelInfos, err
}
