package main

import (
	"fmt"
	"gomod-demo/logagent/kafka"
)

//logAgent入口程序
func main() {
	//1.初始化kafka连接
	err := kafka.Init([]string{"localhost:9092"})
	if err != nil {
		fmt.Printf("init kafka failed, err: %v\n", err)
		return
	}

	//2.打开日志文件准备收集日志
}
