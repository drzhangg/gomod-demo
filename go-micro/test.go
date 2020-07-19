package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
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
	res, err := callApi(node.Address, "/v1/prods", "GET")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func callApi(addr, path, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	return string(data), nil
}
