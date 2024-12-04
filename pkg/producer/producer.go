package producer

import (
	"ctx-interview/queue"
	"encoding/json"
	"os"
)

type Task struct {
	Name    string            `json:"name"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type Producer struct {
	queue *queue.Queue
}

func NewProducer(queue *queue.Queue) *Producer {
	return &Producer{queue: queue}
}

func (p *Producer) LoadTasks(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks struct {
		Tasks []Task `json:"tasks"`
	}
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks.Tasks, nil
}

func (p *Producer) SendTasks(queueName string, tasks []Task) error {
	for _, task := range tasks {
		taskJSON, _ := json.Marshal(task)
		if err := p.queue.Client.Publish(p.queue.Ctx, queueName, string(taskJSON)).Err(); err != nil {
			return err
		}
	}
	return nil
}
