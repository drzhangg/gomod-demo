package main

import (
	"gomod-demo/grpc/stream/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	PORT = "9002"
)

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)

	err = printLists(client, &pb.StreamReq{Pt: &pb.StreamPoint{
		Name:  "gRPC Stream Client: List",
		Value: 2020,
	}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	err = printRecord(client, &pb.StreamReq{Pt: &pb.StreamPoint{
		Name:  "gRPC Stream Client: Record",
		Value: 2020,
	}})
	if err != nil {
		log.Fatalf("printRecord.err: %v", err)
	}

	err = printRoute(client, &pb.StreamReq{Pt: &pb.StreamPoint{
		Name:  "gRPC Stream Client: Route",
		Value: 2020,
	}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}

func printLists(client pb.StreamServiceClient, r *pb.StreamReq) error {
	return nil
}

func printRecord(client pb.StreamServiceClient, r *pb.StreamReq) error {
	return nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamReq) error {
	return nil
}
