package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	wg  sync.WaitGroup
)

func main() {
	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	consumer, err := sarama.NewConsumer([]string{"47.103.9.218:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionList, err := consumer.Partitions("kafka_go_test")
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("kafka_go_test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}

		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
