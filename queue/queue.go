package queue

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Queue struct {
	client *redis.Client
	ctx    context.Context
}

func NewQueue(redisAddr string) *Queue {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &Queue{
		client: client,
		ctx:    context.Background(),
	}
}

func (q *Queue) Enqueue(queueName string, message string) error {
	return q.client.LPush(q.ctx, queueName, message).Err()
}

func (q *Queue) Dequeue(queueName string) (string, error) {
	return q.client.RPop(q.ctx, queueName).Result()
}
