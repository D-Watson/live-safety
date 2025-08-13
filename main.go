package main

import (
	"context"

	"live_safety/conf"
	"live_safety/controller"
	"live_safety/dbs"
	"live_safety/rpc"
	"live_safety/services"
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
	services.InitCron(ctx)
}

func main() {
	ctx := context.Background()
	Init(ctx)
	// 开启定时任务
	// 启动gRPC服务
	go rpc.RunRpcServer()
	controller.InitRouters()

}
