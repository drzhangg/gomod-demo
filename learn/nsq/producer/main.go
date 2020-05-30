package main

import (
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

type producer struct {
	producer *nsq.Producer
}

func (p *producer) publish(topic string, message string) (err error) {
	if message == "" {
		return errors.New("message is empty")
	}
	if err = p.producer.Publish(topic, []byte(message)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//延迟消息
func (p *producer) deferredPublish(topic string, delay time.Duration, message string) (err error) {
	if message == "" {
		return errors.New("message is empty")
	}
	if err = p.producer.DeferredPublish(topic, delay, []byte(message)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InitProducer(addr string) (p *nsq.Producer, err error) {
	var (
		config *nsq.Config
	)
	config = nsq.NewConfig()
	if p, err = nsq.NewProducer(addr, config); err != nil {
		return nil, err
	}
	return p, nil
}

func main() {

}
