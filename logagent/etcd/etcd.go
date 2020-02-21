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

// 用来监听kafka里面的配置文件是否要变动
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	watchCh := client.Watch(context.Background(), key)
	//从通道尝试取值（etcd监视的信息）
	for v := range watchCh {
		for _, ev := range v.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))

			var newConf []*LogEntry
			// 判断操作类型，如果是删除配置操作，则返回空的结构体
			if ev.Type != clientv3.EventTypeDelete {
				if err := json.Unmarshal(ev.Kv.Value, &newConf); err != nil {
					fmt.Printf("unmarshal failed, err:%v\n", err)
					continue
				}
			}

			fmt.Printf(" get new conf:%v\n", newConf)
			// 取得了信息的变化，通过一个通道把变化的信息发送出去
			newConfCh <- newConf
		}
	}
}
