package main

import (
	"context"
	"gomod-demo/grpc/search/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchClient(conn)

	resp, err := client.Search(context.TODO(), &pb.SearchReq{Request: "grpc"})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp:%s", resp.GetResponse())

}
