package rpc

import (
	"context"

	pb "live_safty/proto"
	"live_safty/services"
)

type SafeTransferServer struct {
	pb.UnimplementedTransferSafeServer // 必须内嵌
}

func (s *SafeTransferServer) SecureTransferPublicKey(ctx context.Context, req *pb.GetPublicTokenRequest) (*pb.GetPublicTokenReply, error) {
	// 这里写你的安全传输逻辑！

	return &pb.GetPublicTokenReply{}, nil
}

func (s *SafeTransferServer) SecureTransferPrivateKey(ctx context.Context, req *pb.GetPrivateTokenRequest) (*pb.GetPrivateTokenReply, error) {

	return &pb.GetPrivateTokenReply{}, nil
}

func (s *SafeTransferServer) SecureEncrypt(ctx context.Context, req *pb.Data) (*pb.Data, error) {
	res, err := services.AcquireEncrypt(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SafeTransferServer) SecureDecrypt(ctx context.Context, req *pb.Data) (*pb.Data, error) {
	res, err := services.AcquireDecrypt(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
