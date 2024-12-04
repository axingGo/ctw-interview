package cmd

import (
	"ctx-interview/conf"
	"ctx-interview/pkg/consumer"
	"ctx-interview/pkg/scraper"
	"ctx-interview/pkg/storage"
	"ctx-interview/queue"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var consumerCmd = &cobra.Command{
	Use:   "start-consumer",
	Short: "Start a consumer to process tasks from the queue.",
	Run: func(cmd *cobra.Command, args []string) {
		queueName, _ := cmd.Flags().GetString("queue")
		workers, _ := cmd.Flags().GetInt("workers")

		// 初始化
		q := queue.NewQueue(conf.Conf.Redis.Host)
		s := scraper.NewScraper()
		mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conf.Conf.MySQL.User, conf.Conf.MySQL.Password, conf.Conf.MySQL.Host, conf.Conf.MySQL.Port, conf.Conf.MySQL.DBName)
		db, err := storage.NewDatabase(mysqlDsn)
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
