package dbs

import (
	"context"
	"encoding/json"
	"time"

	"github.com/D-Watson/live-safety/consts"
	"github.com/D-Watson/live-safety/log"
)

type KeyPair struct {
	PrivateKey string `json:"privateKey"` // PEM格式
	PublicKey  string `json:"publicKey"`
	KeyID      string `json:"keyId"` // 唯一标识符
}

func HSetTokenPairsByKey(ctx context.Context, key string, en KeyPair) error {
	ct, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	log.Infof(ctx, "[redis] set key=%s, v=%v", key, en)
	byteData, _ := json.Marshal(en)
	err := GlobalRedisCli.Set(ct, key, byteData, consts.EXPIRE_TIME).Err()
	if err != nil {
		log.Errorf(ctx, "[redis] set error, err=", err)
		return err
	}
	return nil
}

func HGetTokenPairsByKey(ctx context.Context, key string, en *KeyPair) error {
	ct, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	marshal, err := GlobalRedisCli.Get(ct, key).Result()
	if err != nil {
		log.Errorf(ctx, "[redis] get error, err=", err)
		return err
	}
	err = json.Unmarshal([]byte(marshal), &en)
	log.Infof(ctx, "[redis] get key=%s, value=%v", key, en)
	if err != nil {
		return err
	}
	return nil
}
