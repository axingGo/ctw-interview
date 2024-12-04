package consumer

import (
	"ctx-interview/pkg/scraper"
	"ctx-interview/pkg/storage"
	"ctx-interview/queue"
	"encoding/json"
	"log"
	"time"
)

type Consumer struct {
	queue    *queue.Queue
	scraper  *scraper.Scraper
	database *storage.Database
}

func NewConsumer(queue *queue.Queue, scraper *scraper.Scraper, db *storage.Database) *Consumer {
	return &Consumer{
		queue:    queue,
		scraper:  scraper,
		database: db,
	}
}

func (c *Consumer) Start(queueName string) {
	for {
		taskJSON, err := c.queue.Dequeue(queueName)
		if err != nil {
			log.Println("Failed to dequeue task:", err)
			time.Sleep(time.Second)
			continue
		}

		var task scraper.Task
		if err := json.Unmarshal([]byte(taskJSON), &task); err != nil {
			log.Println("Failed to parse task:", err)
			continue
		}

		hotelInfos, err := c.scraper.Scrape(task)
		if err != nil {
			log.Println("Failed to scrape task:", err)
			continue
		}

		if err := c.database.SaveResults(queueName, hotelInfos); err != nil {
			log.Println("Failed to save results:", err)
		}
	}
}
