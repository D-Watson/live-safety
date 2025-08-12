package main

import (
	"context"

	"live_safty/conf"
	"live_safty/dbs"
	"live_safty/rpc"
	"live_safty/services"
)

func Init(ctx context.Context) {
	//1. 解析配置文件
	err := conf.ParseConfig(ctx)
	if err != nil {
		return
	}
	//2. 初始化redis连接
	err = dbs.InitRedisCli(ctx)
	if err != nil {
		return
	}
}

func main() {
	ctx := context.Background()
	Init(ctx)
	// 开启定时任务
	services.InitCron(ctx)
	rpc.RunRpcServer()
}
