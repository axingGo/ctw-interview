package queue

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Queue struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewQueue(redisAddr string) *Queue {
	fmt.Println(redisAddr)
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &Queue{
		Client: client,
		Ctx:    context.Background(),
	}
}
