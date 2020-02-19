package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	client *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string, timeout int) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: time.Duration(timeout) * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd failed, err:", err)
		return err
	}
	return nil
}

func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	getResp, err := client.Get(ctx, key)
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}

	for _, ev := range getResp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Println("unmarshal etcd value failed, err:", err)
			return nil, err
		}
	}
	return logEntryConf, nil
}
