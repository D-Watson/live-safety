package services

import (
	"context"
	"sync"

	"github.com/robfig/cron/v3"
	"live_safety/consts"
	"live_safety/dbs"
	"live_safety/log"
)

func InitCron(ctx context.Context) {
	c := cron.New()
	log.Infof(ctx, "[cron] job started")
	//启动时候执行一次
	cronJob(ctx)
	//每30分钟执行一次
	_, err := c.AddFunc("0 30 * * * *", func() {
		cronJob(ctx)
	})
	if err != nil {
		return
	}
	c.Start()
	defer c.Stop()
	select {}
}

// 定时任务更新密钥对

func cronJob(ctx context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := generateAndSetTokens(ctx, consts.LIVE_TOKEN_BACK_END)
		if err != nil {
			//log.Errorf(ctx, "[generate] token error,err=", err)
			return
		}
	}()
	go func() {
		defer wg.Done()
		err := generateAndSetTokens(ctx, consts.LIVE_TOKEN_FRONT_END)
		if err != nil {
			//log.Errorf(ctx, "[generate] token error,err=", err)
			return
		}
	}()
	wg.Wait()
}

func generateAndSetTokens(ctx context.Context, key string) error {
	prvkey, pubkey, err := GenRsaKey(consts.BIT_SIZE)
	if err != nil {
		log.Errorf(ctx, "[token]generate error , err=", err)
		return err
	}
	kp := dbs.KeyPair{
		KeyID:      key,
		PrivateKey: string(prvkey),
		PublicKey:  string(pubkey),
	}
	err = dbs.HSetTokenPairsByKey(ctx, key, kp)
	if err != nil {
		return err
	}
	return nil
}
