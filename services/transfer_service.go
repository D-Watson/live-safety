package services

import (
	"context"
	"encoding/hex"

	"github.com/D-Watson/live-safety/consts"
	"github.com/D-Watson/live-safety/dbs"
	"github.com/D-Watson/live-safety/entity"
	"github.com/D-Watson/live-safety/log"
	pb "github.com/D-Watson/live-safety/proto"
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

func VerifyTransferParams(ctx context.Context, req *entity.TransferRequest) bool {
	if len(req.TransferData) == 0 {
		log.Errorf(ctx, "[param err] transfer data is nil.")
		return false
	}
	if req.Role != consts.LIVE_BACKEND_REQ && req.Role != consts.LIVE_FRONTEND_REQ {
		log.Errorf(ctx, "[param err] role number is wrong.")
		return false
	}
	if req.Crypto != consts.LIVE_ENCRYPT && req.Crypto != consts.LIVE_DECRYPT {
		log.Errorf(ctx, "[param err] the request is not encrypt or decrypt.")
		return false
	}
	return true
}

func TransferHttp(ctx context.Context, req *entity.TransferRequest) (*entity.TransferResponse, int) {
	kp, err := getTokenByRole(ctx, req.Role)
	resp := &entity.TransferResponse{}
	if err != nil || kp == nil {
		return nil, consts.GET_KEY_ERR
	}
	switch req.Crypto {
	case consts.LIVE_ENCRYPT:
		res, er := RsaEncrypt([]byte(kp.PublicKey), []byte(req.TransferData))
		if er != nil || res == nil {
			return nil, consts.ENCODE_ERR
		}
		resp.TransferData = hex.EncodeToString(res)
		break
	case consts.LIVE_DECRYPT:
		data, er := hex.DecodeString(req.TransferData)
		res, er := RsaDecrypt([]byte(kp.PrivateKey), data)
		if er != nil || res == nil {
			return nil, consts.DECODE_ERR
		}
		resp.TransferData = string(res)
		break
	}
	return resp, 0
}

// AcquireEncrypt 服务端用前端的公钥加密, 前端用服务端的公钥加密
func AcquireEncrypt(ctx context.Context, req *pb.Data) (*pb.Data, error) {
	kp, err := getTokenByRole(ctx, req.Role)
	if err != nil || kp == nil {
		return nil, err
	}
	data := req
	res, err := RsaEncrypt([]byte(kp.PublicKey), []byte(data.TransData))
	if err != nil || res == nil {
		return nil, err
	}
	data.EncryptData = hex.EncodeToString(res)
	return data, nil
}

// AcquireDecrypt 前端用自己的秘钥解密， 服务端用自己的秘钥解密
func AcquireDecrypt(ctx context.Context, req *pb.Data) (*pb.Data, error) {
	kp, err := getTokenByRole(ctx, req.Role)
	if err != nil || kp == nil {
		return nil, err
	}
	resp := req
	data, _ := hex.DecodeString(req.TransData)
	res, err := RsaDecrypt([]byte(kp.PrivateKey), data)
	if err != nil || res == nil {
		return nil, err
	}
	resp.DecryptData = string(res)
	return resp, nil
}
