package config

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisConnector interface {
	Open(conf YamlConfig) (RedisClient, error)
}

type redisConnector struct{}

type RedisClient struct {
	RedisConnection *redis.Client
	Ctx             context.Context
}

func (redisConnector redisConnector) Open(conf YamlConfig) (RedisClient, error) {
	ctx := context.Background()

	db, err := strconv.Atoi(conf.Redis.Db)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return RedisClient{}, nil
	}

	redisConnection := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host + ":" + conf.Redis.Port,
		Username: conf.Redis.User,
		Password: conf.Redis.Pass,
		DB:       db,
	})
	return RedisClient{
		RedisConnection: redisConnection,
		Ctx:             ctx,
	}, nil
}

func NewRedisConnector() redisConnector {
	return redisConnector{}
}
