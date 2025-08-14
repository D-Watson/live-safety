package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/D-Watson/live-safety/conf"
	"github.com/D-Watson/live-safety/dbs"
	"github.com/D-Watson/live-safety/services"
)

func TestRsaEncrypt(t *testing.T) {
	ctx := context.Background()
	conf.ParseConfig(ctx)
	
	services.InitCron(ctx)
	s := "hello223234"
	m, _ := dbs.GlobalRedisCli.Get(context.Background(), "live:token:frontend").Result()
	ke := &dbs.KeyPair{}
	_ = json.Unmarshal([]byte(m), ke)
	fmt.Println(ke)
	res, _ := services.RsaEncrypt([]byte(ke.PublicKey), []byte(s))
	fmt.Println(res)
}
