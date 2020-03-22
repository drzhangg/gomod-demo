package main

import (
	"context"
	"fmt"
	pb "gomod-demo/grpc/pb"
	"google.golang.org/grpc"
	"net"
)

//服务端
/*
1.开启端口
2.实现grpc接口
3.注册grpc
*/

type GetUserReq struct {
}

//var

func GetUserInfo(ctx context.Context, req *pb.GetUserReq) (rsp *pb.GetUserRsp, err error) {
	name := req.Name

	if name != "" {
		rsp = &pb.GetUserRsp{
			Id:    1,
			Name:  name,
			Age:   24,
			Hobby: []string{"动漫", "游戏"},
		}
		return
	}
	return
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("监听异常：%s\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s,&GetUserReq{})
}
