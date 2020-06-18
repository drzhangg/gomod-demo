package main

import (
	"context"
	"fmt"
	"gomod-demo/grpc/cache/pb"
	"google.golang.org/grpc"
	"net"
	"os"
)

type CacheService struct {
}

func (c *CacheService) Get(ctx context.Context, in *pb.GetReq) (*pb.GetResp, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (c *CacheService) Store(ctx context.Context, in *pb.StoreReq) (*pb.StoreResp, error) {
	return nil, fmt.Errorf("unimplemented")
}

func main() {
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run cache server: %s\n", err)
		os.Exit(1)
	}
}

func runServer() error {
	server := grpc.NewServer()
	//注册grpc客户端
	pb.RegisterCacheServer(server, &CacheService{})
	l, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		return err
	}
	return server.Serve(l)
}
