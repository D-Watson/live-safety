package services

import (
	"context"

	pb "live_safty/proto"
)

type safeTransferServer struct {
	pb.UnimplementedTransferSafeServer // 必须内嵌
}

func (s *safeTransferServer) SecureTransferPublicKey(ctx context.Context, req *pb.GetPublicTokenRequest) (*pb.GetPublicTokenReply, error) {
	// 这里写你的安全传输逻辑！

	return &pb.GetPublicTokenReply{}, nil
}
