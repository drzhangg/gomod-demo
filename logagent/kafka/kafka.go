package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func Init(addr []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition分区
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel中返回

	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return err
	}
	return nil
}

func SentToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err :", err)
		return
	}
	fmt.Println(msg.Value)
	fmt.Println("pid:", pid, "offset:", offset)
}
