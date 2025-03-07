package adapter

import (
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/revandpratama/reflect/config"
	"github.com/revandpratama/reflect/helper"
)

type RedisOption struct {
	RedisClient *redis.Client
}

func (r *RedisOption) Start(a *Adapter) error {
	helper.NewLog().Info("initializing redis...").ToKafka()

	redisDB, err := strconv.Atoi(config.ENV.RedisDB)
	if err != nil {
		return err
	}

	r.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.ENV.RedisHost, config.ENV.RedisPort),
		Password: config.ENV.RedisPassword,
		DB:       redisDB,
	})
	a.RedisClient = r.RedisClient

	helper.NewLog().Info("redis running").ToKafka()
	return nil
}

func (r *RedisOption) Stop() error {
	return r.RedisClient.Close()
}
