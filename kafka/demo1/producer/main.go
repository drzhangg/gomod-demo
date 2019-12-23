package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()

	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll

	//随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	//是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	config.Version = sarama.V0_11_0_2

	//使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewAsyncProducer([]string{"47.103.9.218:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.AsyncClose()

	//构建要发送的消息
	msg := &sarama.ProducerMessage{
		Topic:     "hello",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var value string
	for {
		_, err := fmt.Scanf("%s", &value)
		if err != nil {
			break
		}
		msg.Value = sarama.ByteEncoder(value)
		fmt.Println(value)

		producer.Input() <- msg
		//partition, offset, err := producer.SendMessage(msg)
		//if err != nil {
		//	fmt.Println("Send message Fail")
		//}
		//fmt.Printf("Partition = %d, offset=%d\n", partition, offset)

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}
