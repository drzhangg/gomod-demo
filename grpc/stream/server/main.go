package main

import (
	"gomod-demo/grpc/stream/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type StreamService struct {
}

const (
	PORT = "9002"
)

func (s *StreamService) List(req *pb.StreamReq, stream pb.StreamService_ListServer) error {
	for i := 0; i <= 6; i++ {
		err := stream.Send(&pb.StreamResp{
			Pt: &pb.StreamPoint{
				Name:  req.Pt.Name,
				Value: req.Pt.Value,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResp{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	return nil
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResp{Pt: &pb.StreamPoint{
			Name:  "gPRC Stream Client: Route",
			Value: int32(n),
		}})
		if err != nil {
			return err
		}
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

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
