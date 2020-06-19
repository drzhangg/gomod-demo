package main

import (
	"context"
	"fmt"
	pb2 "gomod-demo/grpc/user/pb"
	"google.golang.org/grpc"
	"net"
)

//服务端
/*
1.开启端口
2.实现grpc接口
3.注册grpc
*/

type UserServer struct {
}

//var u = UserServer{}

func (u *UserServer) GetUserInfo(ctx context.Context, req *pb2.GetUserReq) (rsp *pb2.GetUserRsp, err error) {
	name := req.Name

	if name != "" {
		rsp = &pb2.GetUserRsp{
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

	fmt.Println("正在监听grpc服务...")
	s := grpc.NewServer()

	//注册grpc服务
	pb2.RegisterUserServer(s, &UserServer{})

	s.Serve(listener)
}
