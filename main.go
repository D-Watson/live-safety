package main

import (
	"context"

	"github.com/D-Watson/live-safety/conf"
	"github.com/D-Watson/live-safety/controller"
	"github.com/D-Watson/live-safety/dbs"
	"github.com/D-Watson/live-safety/rpc"
	"github.com/D-Watson/live-safety/services"
)

func Init(ctx context.Context) {
	//1. 解析配置文件
	err := conf.ParseConfig(ctx)
	if err != nil {
		return
	}
	//2. 初始化redis连接
	err = dbs.InitRedis(ctx)
	if err != nil {
		return
	}
	services.InitCron(ctx)
}

func main() {

	ctx := context.Background()
	Init(ctx)
	//开启定时任务
	//启动gRPC服务
	go rpc.RunRpcServer()
	controller.InitRouters()
}
