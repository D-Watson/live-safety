package dbs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"live_safety/conf"
)

var GlobalRedisCli *redis.Client

func InitRedisCli(ctx context.Context) error {
	GlobalRedisCli = redis.NewClient(&redis.Options{
		Addr:     conf.GlobalConfig.DB.Address, // Redis地址
		Password: conf.GlobalConfig.DB.Passwd,  // 密码（无密码留空）
		DB:       0,                            // 默认DB
	})
	// 测试连接
	pong, err := GlobalRedisCli.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println("连接成功:", pong)
	return nil
}
