package main

import (
	"fmt"
	"gomod-demo/logagent/conf"
	"gomod-demo/logagent/kafka"
	"gomod-demo/logagent/taillog"
	"gopkg.in/ini.v1"
)

var (
	config = new(conf.Config)
)

func run() {
	//1.从tail读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2.把日志写入kafka
			kafka.SentToKafka(config.Topic, line.Text)
		default:

		}
	}
}

//logAgent入口程序
func main() {
	//加载配置文件
	err := ini.MapTo(config, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini config file failed, err：%v\n", err)
		return
	}
	fmt.Println(config.Address)
	fmt.Println(config.Topic)
	fmt.Println(config.Path)

	//1.初始化kafka连接
	err = kafka.Init([]string{config.Address})
	if err != nil {
		fmt.Printf("init kafka failed, err: %v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	//2.打开日志文件准备收集日志
	err = taillog.Init(config.Path)
	if err != nil {
		fmt.Printf("init tail log failed,err: %v\n", err)
		return
	}
	fmt.Println("init taillog success.")

	run()
}
