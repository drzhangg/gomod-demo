package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"time"
)

type MyTestHandler struct {
	consumer *nsq.Consumer
}

func main() {
	adds := []string{"127.0.0.1:7201", "127.0.0.1:8201"}
	config := nsq.NewConfig()
	config.MaxInFlight = 1000
	config.MaxBackoffDuration = 5 * time.Second
	config.DialTimeout = 10 * time.Second

	topicName := "testTopic1"
	consumer, err := nsq.NewConsumer(topicName, "ch1", config)
	if err != nil {
		fmt.Println(err)
	}
	testHandler := &MyTestHandler{consumer: consumer}
	consumer.AddHandler(testHandler)

	err = consumer.ConnectToNSQLookupds(adds)
	if err != nil {
		panic(err)
	}
	status := consumer.Stats()
	if status.Connections == 0 {
		panic("stats report 0 connections (should be > 0)")
	}
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("server is running.....")
	<-stop

}

func (m MyTestHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	return nil
}
