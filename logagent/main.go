package main

import (
	"fmt"
	"gomod-demo/logagent/kafka"
	"gomod-demo/logagent/taillog"
)

func run() {
	//1.从tail读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2.把日志写入kafka
			kafka.SentToKafka("web_log", line.Text)
		default:

		}
	}
}

//logAgent入口程序
func main() {
	//1.初始化kafka连接
	err := kafka.Init([]string{"localhost:9092"})
	if err != nil {
		fmt.Printf("init kafka failed, err: %v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	//2.打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("init tail log failed,err: %v\n", err)
		return
	}
	fmt.Println("init taillog success.")

	run()
}
