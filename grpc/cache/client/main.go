package main

import (
	"context"
	"fmt"
	"gomod-demo/grpc/cache/pb"
	"google.golang.org/grpc"
	"os"
)

func main() {
	if err := runClient(); err != nil {
		fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		os.Exit(1)
	}
}

func runClient() error {
	//简历grpc连接
	conn, err := grpc.Dial("localhost:5053", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial server: %v", err)
	}

	//初始化grpc客户端
	cacheClient := pb.NewCacheClient(conn)

	//调用grpc客户端方法
	_, err = cacheClient.Store(context.TODO(), &pb.StoreReq{
		Key: "name",
		Val: []byte("jerry"),
	})
	if err != nil {
		return fmt.Errorf("failed to store: %v", err)
	}

	resp, err := cacheClient.Get(context.TODO(), &pb.GetReq{Key: "jerry"})
	if err != nil {
		return fmt.Errorf("failed to get: %v", err)
	}

	fmt.Printf("Get cache value: %s\n", resp.Val)
	return nil
}
