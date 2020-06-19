package main

import (
	"gomod-demo/grpc/stream/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type StreamService struct {
}

const (
	PORT = "9002"
)

func (s *StreamService) List(req *pb.StreamReq, stream pb.StreamService_ListServer) error {
	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	return nil
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	return nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
