package services

import (
	"context"

	"github.com/robfig/cron/v3"
	"live_safty/consts"
	"live_safty/dbs"
	"live_safty/log"
)

func InitCron(ctx context.Context) {
	c := cron.New()
	//每30分钟执行一次
	_, err := c.AddFunc("0 30 * * * *", func() {
		err := generateAndSetTokens(ctx, consts.LIVE_TOKEN_FRONT_END)
		if err != nil {
			return
		}
		err = generateAndSetTokens(ctx, consts.LIVE_TOKEN_BACK_END)
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	c.Start()
	defer c.Stop()
	select {}
}

// 定时任务更新密钥对
func generateAndSetTokens(ctx context.Context, key string) error {
	prvkey, pubkey, err := GenRsaKey(consts.BIT_SIZE)
	if err != nil {
		log.Errorf(ctx, "[token]generate error , err=", err)
		return err
	}
	kp := &dbs.KeyPair{
		KeyID:      key,
		PrivateKey: prvkey,
		PublicKey:  pubkey,
	}
	err = dbs.HSetTokenPairsByKey(ctx, key, kp)
	if err != nil {
		return err
	}
	return nil
}
