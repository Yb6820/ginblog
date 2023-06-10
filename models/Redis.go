package models

import (
	"context"
	"fmt"
	"ginblog/utils"

	"github.com/redis/go-redis/v9"
)

var Red *redis.Client

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         utils.RedisAddr,
		Password:     utils.RedisPassword,
		DB:           utils.DB,
		PoolSize:     utils.PoolSize,
		MinIdleConns: utils.MinIdleConn,
	})
	ctx, cancel := context.WithCancel(context.Background())

	pong, err := Red.Ping(ctx).Result()
	if err != nil {
		fmt.Println("init redis ....", err)
	} else {
		fmt.Println("Redis inited ....", pong)
	}
	cancel()
}
