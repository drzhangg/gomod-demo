package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type messageHandler struct {
}

func (m *messageHandler) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		return nil
	}

	log.Printf("Got a message:%s", message.Body)
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("text", "channel", config)
	if err != nil {
		log.Fatalf("nsq consumer failed, err:%v\n", err)
	}

	consumer.AddHandler(&messageHandler{})

	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Fatal(err)
	}
	consumer.Stop()
}
