package cmd

import (
	"ctx-interview/pkg/producer"
	"ctx-interview/queue"
	"log"

	"github.com/spf13/cobra"
)

var producerCmd = &cobra.Command{
	Use:   "start producer",
	Short: "Start a producer to send tasks to the queue.",
	Run: func(cmd *cobra.Command, args []string) {
		queueName, _ := cmd.Flags().GetString("queue")
		taskFile, _ := cmd.Flags().GetString("data")

		// 初始化队列
		q := queue.NewQueue("127.0.0.1:6379")
		p := producer.NewProducer(q)

		// 加载任务并发送到队列
		tasks, err := p.LoadTasks(taskFile)
		if err != nil {
			log.Fatalf("Failed to load tasks: %v", err)
		}

		if err := p.SendTasks(queueName, tasks); err != nil {
			log.Fatalf("Failed to send tasks: %v", err)
		}

		log.Println("Tasks have been successfully sent to the queue.")
	},
}

func init() {
	producerCmd.Flags().StringP("queue", "q", "airbnb_queue", "The name of the queue to use")
	producerCmd.Flags().StringP("data", "d", "tasks.json", "Path to the tasks JSON file")
	RootCmd.AddCommand(producerCmd)
}
