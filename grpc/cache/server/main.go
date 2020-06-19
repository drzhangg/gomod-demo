package main

import (
	"context"
	"fmt"
	"gomod-demo/grpc/cache/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"os"
)

type CacheService struct {
	store map[string][]byte
}

func (c *CacheService) Get(ctx context.Context, in *pb.GetReq) (*pb.GetResp, error) {
	val, ok := c.store[in.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key not found %s", in.Key)
	}
	return &pb.GetResp{Val: val}, nil
}

func (c *CacheService) Store(ctx context.Context, in *pb.StoreReq) (*pb.StoreResp, error) {
	c.store = make(map[string][]byte)
	c.store[in.Key] = in.Val
	return &pb.StoreResp{}, nil
}

func main() {
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run cache server: %s\n", err)
		os.Exit(1)
	}
}

func runServer() error {
	server := grpc.NewServer()

	cache := new(CacheService)
	//注册grpc客户端
	pb.RegisterCacheServer(server, cache)
	l, err := net.Listen("tcp", "localhost:5053")
	if err != nil {
		return err
	}
	return server.Serve(l)
}
