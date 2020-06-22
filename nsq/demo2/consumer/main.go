package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1000)

	config := nsq.NewConfig()
	c, _ := nsq.NewConsumer("testTopic", "ch", config)
	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Get a message: %s", message.Body)
		wg.Done()
		return nil
	}))

	err := c.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Panic(err)
	}
	wg.Wait()
}
