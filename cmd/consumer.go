package cmd

import (
	"ctx-interview/pkg/consumer"
	"ctx-interview/pkg/scraper"
	"ctx-interview/pkg/storage"
	"ctx-interview/queue"
	"log"

	"github.com/spf13/cobra"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start a consumer to process tasks from the queue.",
	Run: func(cmd *cobra.Command, args []string) {
		queueName, _ := cmd.Flags().GetString("queue")
		workers, _ := cmd.Flags().GetInt("workers")

		// 初始化
		q := queue.NewQueue("127.0.0.1:6379")
		s := scraper.NewScraper()
		db, err := storage.NewDatabase("user:password@tcp(localhost:3306)/airbnb")
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
			panic(err)
		}

		// 启动消费者
		c := consumer.NewConsumer(q, s, db)
		for i := 0; i < workers; i++ {
			go c.Start(queueName)
		}

		log.Printf("Consumer started with %d workers. Listening to queue: %s", workers, queueName)
		select {}
	},
}

func init() {
	consumerCmd.Flags().StringP("queue", "q", "airbnb_queue", "The name of the queue to listen to")
	consumerCmd.Flags().IntP("workers", "w", 5, "Number of concurrent workers")
	RootCmd.AddCommand(consumerCmd)
}
