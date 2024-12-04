package consumer

import (
	"ctx-interview/pkg/scraper"
	"ctx-interview/pkg/storage"
	"ctx-interview/queue"
	"encoding/json"
	"fmt"
	"log"
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
	// 订阅任务 Channel
	sub := c.queue.Client.Subscribe(c.queue.Ctx, queueName)
	defer sub.Close()

	fmt.Println("Subscribed to channel:", queueName)

	for {
		msg, err := sub.ReceiveMessage(c.queue.Ctx)
		if err != nil {
			fmt.Println("Error receiving message:", err)
			break
		}

		var task scraper.Task
		if err := json.Unmarshal([]byte(msg.Payload), &task); err != nil {
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
