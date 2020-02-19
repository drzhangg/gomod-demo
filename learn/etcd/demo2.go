package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd failed, err:", err)
		return
	}
	fmt.Println("connect etcd success")
	defer client.Close()

	watchChan := client.Watch(context.Background(), "testkey", clientv3.WithPrefix())
	for wch := range watchChan {
		for _, evt := range wch.Events {
			fmt.Println(evt.Type)
			fmt.Println(string(evt.Kv.Key))
			fmt.Println(string(evt.Kv.Value))
		}
	}

}
