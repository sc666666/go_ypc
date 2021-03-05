package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go_ypc/pkg/config"
	"time"
)

var (
	RDB *redis.Client
	CTX = context.Background()
)

func InitClient() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.GetString("database.redis.default.host") + ":" + config.GetString("database.redis.default.port"),
		Password: config.GetString("database.redis.default.password"),
		DB:       config.GetInt("database.redis.default.database"),
	})

	ctx, cancel := context.WithTimeout(CTX, 5*time.Second)
	defer cancel()

	_, err = RDB.Ping(ctx).Result()
	return err
}
