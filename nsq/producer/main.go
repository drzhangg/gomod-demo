package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

func main() {
	//定义nsq生产者
	var producer *nsq.Producer

	//初始化生产者
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	err = producer.Ping()
	if nil != err {
		// 关闭生产者
		producer.Stop()
		producer = nil
	}

	//生产者写入nsq，10条消息，topic="test"
	topic := "test"
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("message:%d", i)
		if producer != nil && message != "" { //不能发布空字符串，否则会导致error
			err = producer.Publish(topic, []byte(message))
			if err != nil {
				fmt.Printf("producer.Publish,err : %v", err)
			}
			fmt.Println(message)
		}
	}
	fmt.Println("producer.Publish success")
}
