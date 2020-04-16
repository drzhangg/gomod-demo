package main

import (
	"bufio"
	"context"
	pb "gomod-demo/protobuf/pb/chat"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

func main() {
	//创建连接
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败：[%v]\n", err)
		return
	}
	defer conn.Close()

	//声明客户端
	client := pb.NewChatClient(conn)

	ctx := context.Background()

	//创建双向数据流
	stream, err := client.BidStream(ctx)
	if err != nil {
		log.Printf("创建数据流失败：[%v]\n", err)
	}

	//启动一个 goroutine 接收命令行输入参数
	go func() {
		log.Println("请输入消息：")
		reader := bufio.NewReader(os.Stdin)
		for {
			str, _ := reader.ReadString('\n')

			//向服务端发送指令
			if err := stream.Send(&pb.Request{Input: str}); err != nil {
				return
			}
		}
	}()

	for {
		//接收从 服务端返回的数据流
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("⚠️ 收到服务端的结束信号")
			break //如果收到结束信号，则退出“接收循环”，结束客户端程序
		}

		if err != nil {
			// TODO: 处理接收错误
			log.Println("接收数据出错:", err)
		}

		// 没有错误的情况下，打印来自服务端的消息
		log.Printf("[客户端收到]: %s", response.Output)
	}
}
