package scraper

import (
	"ctx-interview/pkg/storage"
	"fmt"
	"strconv"

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
	if err != nil {
		fmt.Println("Request Visit failed:", err)
	}

	// 到详情页获取hotel具体信息
	var hotelInfos []*storage.HotelInfo
	for _, href := range hrefs {
		// todo
		hotel, err := ScrapeHotelDetail(href)
		if err != nil {
			fmt.Printf("ScrapeHotelDetail url:%s,failed:%v\n", href, err)
			continue
		}
		hotelInfos = append(hotelInfos, hotel)
	}
	return hotelInfos, err
}

func ScrapeHotelDetail(href string) (*storage.HotelInfo, error) {
	var hotel *storage.HotelInfo
	c := colly.NewCollector()
	// 获取酒店名称
	c.OnHTML(`h2.hpipapi`, func(e *colly.HTMLElement) {
		hotel.HotelName = e.Text
	})
	// 获取评分名称
	c.OnHTML(`div[data-testid="pdp-reviews-highlight-banner-host-rating"] div[aria-hidden="true"]`, func(e *colly.HTMLElement) {
		rating := e.Text
		ratingF, _ := strconv.ParseFloat(rating, 64)
		hotel.Star = int(ratingF * 100)
	})
	// todo 获取其他信息

	// 处理请求错误
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request failed:", err)
	})
	err := c.Visit(href)
	if err != nil {
		fmt.Println("Visit failed:", err)
	}
	return hotel, nil
}
