package configs

import (
	"fmt"
	"gin-web/helper"
	"time"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db" json:"db"`
	MaxSize  int    `yaml:"max_size"`
	MinIdle  int    `yaml:"min_idle"`
	Timeout  int    `yaml:"timeout"`
}

var REDIS *redis.Client

func InitRedisConfig(redisConfig RedisConfig) {
	REDIS = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		DB:           redisConfig.Db,
		Password:     redisConfig.Password,
		PoolSize:     redisConfig.MaxSize,
		MinIdleConns: redisConfig.MinIdle,
		PoolTimeout:  time.Duration(redisConfig.Timeout) * time.Second,
	})

	var err = REDIS.Ping().Err()

	if err != nil {
		helper.PanicErrorAndMessage(err, "连接redis失败")
	}
}
