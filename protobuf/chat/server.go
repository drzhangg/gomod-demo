package main

import (
	pb "gomod-demo/protobuf/pb/chat"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type Stream struct {
}

func (s *Stream) BidStream(ser pb.Chat_BidStreamServer) error {
	ctx := ser.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("客户端发送的数据流结束")
			return ctx.Err()
		default:
			//接收从客户端发来的消息
			request, err := ser.Recv()
			if err != nil {
				log.Println("接收数据出错：", err)
				return err
			}
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}

			//如果接收正常，则根据接收到的字符串执行相应的指令
			switch request.Input {
			case "结束对话\n":
				log.Println("收到'结束对话'指令")
				if err := ser.Send(&pb.Response{Output: "收到结束指令"}); err != nil {
					return err
				}
				//收到结束指令时，通过return nil终止双向数据流
				return nil

			case "返回数据流\n":
				log.Println("收到'返回数据流'指令")
				for i := 0; i < 10; i++ {
					if err := ser.Send(&pb.Response{Output: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}

			default:
				//缺省情况下，返回 '服务端返回：' + 输入信息
				log.Printf("[收到消息：%s]", request.Input)
				if err := ser.Send(&pb.Response{Output: "服务端返回：" + request.Input}); err != nil {
					return err
				}
			}
		}
	}
}

func main() {
	log.Println("启动服务端....")
	server := grpc.NewServer()

	//注册ChatServer

	pb.RegisterChatServer(server, &Stream{})

	address, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
