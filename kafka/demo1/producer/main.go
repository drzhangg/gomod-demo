package main

import "github.com/Shopify/sarama"

func main() {
	config := sarama.NewConfig()

	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
}
