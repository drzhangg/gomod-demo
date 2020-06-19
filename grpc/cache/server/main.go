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
	accounts      pb.AccountsClient
	store         map[string][]byte
	keysByAccount map[string]int64
}

func (c *CacheService) Get(ctx context.Context, in *pb.GetReq) (*pb.GetResp, error) {
	val, ok := c.store[in.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key not found %s", in.Key)
	}
	return &pb.GetResp{Val: val}, nil
}

func (c *CacheService) Store(ctx context.Context, in *pb.StoreReq) (*pb.StoreResp, error) {
	//调用另一个服务取得账户信息，包含其键值限制
	resp, err := c.accounts.GetByToken(context.TODO(), &pb.GetByTokenReq{Token: in.AccountToken})
	if err != nil {
		return nil, err
	}

	//检查是否超量使用
	if c.keysByAccount[in.AccountToken] >= resp.Account.MaxCacheKeys {
		return nil, status.Errorf(codes.FailedPrecondition, "Account %s exceeds max key limit %d", in.AccountToken, resp.Account.MaxCacheKeys)
	}

	//如果不存在，需要新加键值，对计数器加1
	if _, ok := c.store[in.Key]; !ok {
		c.keysByAccount[in.AccountToken] += 1
	}
	//  保存键值
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
