package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

type ConsumerT struct {
}

func main() {
	InitConsumer("test", "test-channel", "127.0.0.1:4161")
	for {
		time.Sleep(time.Second * 10)
	}
}

func (c *ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

//初始化消费者
func InitConsumer(topic string, channel string, address string) {
	cfg := nsq.NewConfig()

	cfg.LookupdPollInterval = time.Second //设置重连时间

	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		panic(err)
	}

	consumer.SetLogger(nil, 0)        //屏蔽系统日志
	consumer.AddHandler(&ConsumerT{}) // 添加消费者接口

	if err := consumer.ConnectToNSQLookupd(address); err != nil {
		panic(err)
	}
}
