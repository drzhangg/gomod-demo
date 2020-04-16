package main

import (
	"context"
	"fmt"
	pb "gomod-demo/grpc/pb"
	"google.golang.org/grpc"
)

//客户端
/*
1.连接ip端口
2.实例化grpc客户端
3.组装请求参数
*/

func main() {
	//1.
	client, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}

	//2.
	userClient := pb.NewUserClient(client)

	user := pb.GetUserReq{
		Name: "jerry",
	}

	userRsp, err := userClient.GetUserInfo(context.Background(), &user)
	if err != nil {
		fmt.Printf("err：%s\n", err)
	}
	fmt.Println("客户端参数：", userRsp)

}
