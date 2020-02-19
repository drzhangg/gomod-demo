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
		return
	}

	val := `[{"path":"c:/redis/redis.log","topic":"redis_log"},{"path":"c:/mysql/mysql.log","topic":"mysql_log"},{"path":"c:/etcd/etcd.log","topic":"etcd_log"}]`

	_, err = client.Put(context.Background(), "/logagent/collect_config", val)
	if err != nil {
		fmt.Println("etcd put failed:", err)
		return
	}
}
