package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/protoc-gen-micro/examples/greeter"
)

func main() {
	service := micro.NewService(micro.Name("hello"), micro.Version("latest"), micro.Metadata(map[string]string{
		"type": "hello",
	}))

	service.Init()

	greeterServer := greeter.NewGreeterService("hello", service.Client())
	response, err := greeterServer.Hello(context.TODO(), &greeter.Request{Name: "jerry"})
	if err != nil {
		fmt.Println("greeterServer.Hello err:",err)
		return
	}
	fmt.Println(response.String())
}
