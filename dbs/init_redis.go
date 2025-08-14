package dbs

import (
	"context"
	"fmt"

	"github.com/D-Watson/live-safety/conf"
	"github.com/redis/go-redis/v9"
)

var GlobalRedisCli *redis.Client

// InitRedis init redis
func InitRedis(ctx context.Context) error {
	var err error
	GlobalRedisCli, err = InitRedisCli(ctx)
	return err
}

func InitRedisCli(ctx context.Context) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     conf.GlobalConfig.DB.Redis.Address, // Redis地址
		Password: conf.GlobalConfig.DB.Redis.Passwd,  // 密码（无密码留空）
		DB:       0,                                  // 默认DB
	})
	// 测试连接
	pong, err := cli.Ping(ctx).Result()
	if err != nil {
		return cli, err
	}
	fmt.Println("连接成功:", pong)
	return cli, nil
}
