package dbs

import (
	"context"
	"time"

	"live_safty/consts"
	"live_safty/log"
)

type KeyPair struct {
	PrivateKey []byte `json:"privateKey"` // PEM格式
	PublicKey  []byte `json:"publicKey"`
	KeyID      string `json:"keyId"` // 唯一标识符
}

func HSetTokenPairsByKey(ctx context.Context, key string, en *KeyPair) error {
	ct, cancel := context.WithTimeout(ctx, 2)
	defer cancel()
	GlobalRedisCli.Expire(ctx, key, time.Duration(consts.EXPIRE_TIME))
	err := GlobalRedisCli.HSet(ct, key, en).Err()
	if err != nil {
		log.Errorf(ctx, "[redis] hsest error, err=", err)
		return err
	}
	return nil
}

func HGetTokenPairsByKey(ctx context.Context, key string, en *KeyPair) error {
	ct, cancel := context.WithTimeout(ctx, 1)
	defer cancel()
	err := GlobalRedisCli.HGetAll(ct, key).Scan(&en)
	if err != nil {
		log.Errorf(ctx, "[redis] Hgetall error, err=", err)
		return err
	}
	return nil
}
