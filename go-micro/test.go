package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("47.103.9.218:8500"))

	services, err := consulReg.GetService("prodservice1")
	if err != nil {
		log.Fatal(err)
	}

	next := selector.Random(services)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Id, node.Address, node.Metadata)
}
