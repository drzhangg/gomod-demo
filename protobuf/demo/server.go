package main

import (
	"fmt"
	"golang.org/x/net/context"
	pb "gomod-demo/protobuf/pb/customer"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

type server struct {
}

func (s server) GetUser(c context.Context, in *pb.Id) (*pb.User, error) {
	resp := &pb.User{
		Name: fmt.Sprintf("%d,dongtech", in.Id),
		Time: time.Now().UnixNano() / 1e6,
	}
	fmt.Println(resp)
	return resp, nil
}

func (s server) GetActivity(c context.Context, in *pb.Name) (*pb.Activity, error) {
	tp := pb.Tp(rand.Intn(4))
	resp := &pb.Activity{
		Name: in.Name,
		Tp:   tp,
	}
	fmt.Println(resp)
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10086") //grpc服务地址
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, server{})
	log.Println("grpc serve running")
	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
