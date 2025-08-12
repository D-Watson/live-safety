package rpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"live_safty/log"
	"live_safty/proto"
)

func RunRpcServer() {
	ctx := context.Background()
	lis, err := net.Listen("tcp", ":50051")
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
