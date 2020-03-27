package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "gomod-demo/micro/pb"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, in *pb.Request, out *pb.Response) error {
	out.Msg = "hello," + in.Name
	return nil
}

func main() {
	//listener, err := net.Listen("tcp", "127.0.0.1:8080")
	//if err != nil {
	//	fmt.Println("net.Listen err:", err)
	//	return
	//}
	//defer listener.Close()

	service := micro.NewService(
		micro.Name("hello"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type":"hello",
		}))

	service.Init()


	pb.RegisterGreeterHandler(service.Server(),&Greeter{})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	} // 运行服务
}
