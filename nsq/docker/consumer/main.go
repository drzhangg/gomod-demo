package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {
	config :=nsq.NewConfig()
	q,err := nsq.NewConsumer("write_test","ch",config)
	if err !=nil{
		log.Panic("Could not create consumer.")
	}
	defer q.Stop()
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error{
		log.Printf("Got a message: %v",string(message.Body))
		time.Sleep(5*time.Second)
		return nil
	}))
	//err = q.ConnectToNSQD("192.168.9.111:32771");
	err = q.ConnectToNSQD("47.103.9.218:4150");
	if err !=nil {
		log.Panic("Could not connect")
	}
	time.Sleep(3600*time.Second)
}