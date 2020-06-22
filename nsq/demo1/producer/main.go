package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	err = producer.Publish("test", []byte("hello nsq!!!!"))
	if err != nil {
		log.Fatalf("producer publish failed, err:%v\n", err)
	}

	producer.Stop()
}
