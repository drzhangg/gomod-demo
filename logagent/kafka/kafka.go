package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

func Init(addr []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition分区
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel中返回

	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return err
	}
	// 初始化logDataChan
	logDataChan = make(chan *logData, maxSize)
	//开启后台的goroutine从通道中取数据发往kafka
	go SendToKafka()
	return nil
}

// 给外部暴露一个函数，该函数只把数据发送到一个内部的channel中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

// 真正往kafka发送日志的函数
func SendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed,err :", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
