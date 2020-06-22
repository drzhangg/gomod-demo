package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < 1000; i++ {
		msg := fmt.Sprintf("num-%d", i)
		log.Println("Pub:" + msg)

		err = p.Publish("testTopic", []byte(msg))
		if err != nil {
			log.Panic(err)
		}
		time.Sleep(time.Second * 1)
	}

	p.Stop()

}
