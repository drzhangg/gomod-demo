package main

import (
	"context"
	"gomod-demo/grpc/search/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SearchService struct {
}

func (s *SearchService) Search(ctx context.Context, req *pb.SearchReq) (*pb.SearchResp, error) {
	return &pb.SearchResp{Response: req.GetRequest() + " server"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
