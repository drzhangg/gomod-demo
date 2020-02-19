package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	conf := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(conf)
	if err != nil {
		fmt.Println("connect to etcd failed, err:", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//put
	_, err = client.Put(ctx, "testkey", "test value")
	if err != nil {
		fmt.Println("etcd put failed, err:", err)
		return
	}

	//get
	getResp, err := client.Get(ctx, "testkey")
	if err != nil {
		fmt.Println("etcd put failed, err:", err)
		return
	}

	for _, v := range getResp.Kvs {
		fmt.Println("get key:", string(v.Key), " get value:", string(v.Value))
	}

}
