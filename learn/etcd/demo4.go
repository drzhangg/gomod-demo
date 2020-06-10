package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 10,
	}
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Errorf("clientv3.New err:%v", err)
	}

	getResp, err := client.Get(context.TODO(), "name", clientv3.WithPrefix())
	if err != nil {
		fmt.Errorf("clientv3.Get err:%v", err)
	}

	for _, v := range getResp.Kvs {
		fmt.Println(string(v.Value))
	}
}
