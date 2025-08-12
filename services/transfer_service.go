package services

import (
	"context"

	"live_safty/consts"
	"live_safty/dbs"
	"live_safty/log"
	pb "live_safty/proto"
)

func getTokenByRole(ctx context.Context, role int32) (*dbs.KeyPair, error) {
	en := &dbs.KeyPair{}
	if role == consts.LIVE_FRONTEND_REQ {
		err := dbs.HGetTokenPairsByKey(ctx, consts.LIVE_TOKEN_FRONT_END, en)
		if err != nil {
			log.Errorf(ctx, "[service] get token error,key=%s", err, consts.LIVE_TOKEN_FRONT_END)
			return nil, err
		}
	}
	if role == consts.LIVE_BACKEND_REQ {
		err := dbs.HGetTokenPairsByKey(ctx, consts.LIVE_TOKEN_BACK_END, en)
		if err != nil {
			log.Errorf(ctx, "[service] get token error, key = %s,", err, consts.LIVE_TOKEN_BACK_END)
			return nil, err
		}
	}
	return en, nil
}

// AcquireEncrypt 服务端用前端的公钥加密, 前端用服务端的公钥加密
func AcquireEncrypt(ctx context.Context, req *pb.Data) (*pb.Data, error) {
	kp, err := getTokenByRole(ctx, req.Role)
	if err != nil || kp == nil {
		return nil, err
	}
	data := req
	res, err := RsaEncrypt(kp.PublicKey, []byte(data.TransData))
	if err != nil || res == nil {
		return nil, err
	}
	data.EncryptData = string(res)
	return data, nil
}

// AcquireDecrypt 前端用自己的秘钥解密， 服务端用自己的秘钥解密
func AcquireDecrypt(ctx context.Context, req *pb.Data) (*pb.Data, error) {
	kp, err := getTokenByRole(ctx, req.Role)
	if err != nil || kp == nil {
		return nil, err
	}
	data := req
	res, err := RsaDecrypt(kp.PrivateKey, []byte(data.TransData))
	if err != nil || res == nil {
		return nil, err
	}
	data.EncryptData = string(res)
	return data, nil
}
