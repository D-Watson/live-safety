package rpc

import (
	"context"
	"fmt"
	"net"

	"github.com/D-Watson/live-safety/conf"
	"github.com/D-Watson/live-safety/log"
	"github.com/D-Watson/live-safety/proto"
	"google.golang.org/grpc"
)

func RunRpcServer() {
	ctx := context.Background()
	lis, err := net.Listen("tcp", conf.GlobalConfig.Server.Rpc.Host)
	if err != nil {
		log.Errorf(ctx, "failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterTransferSafeServer(s, &SafeTransferServer{})
	fmt.Println("gRPC server is listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf(ctx, "failed to serve: %v", err)
	}
}
