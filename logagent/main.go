package main

import (
	"fmt"
	"gomod-demo/logagent/conf"
	"gomod-demo/logagent/etcd"
	"gomod-demo/logagent/kafka"
	"gomod-demo/logagent/taillog"
	"gopkg.in/ini.v1"
)

var (
	config = new(conf.Config)
)

//func run() {
//	//1.从tail读取日志
//	for {
//		select {
//		case line := <-taillog.ReadChan():
//			//2.把日志写入kafka
//			kafka.SentToKafka(config.Topic, line.Text)
//		default:
//		}
//	}
//}

//logAgent入口程序
func main() {
	//加载配置文件
	err := ini.MapTo(config, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini config file failed, err：%v\n", err)
		return
	}

	//1.初始化kafka连接
	err = kafka.Init([]string{config.Address})
	if err != nil {
		fmt.Printf("init kafka failed, err: %v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	//2. 初始化etcd配置
	err = etcd.Init(config.Endpoints, config.DialTimeout)
	if err != nil {
		fmt.Println("init etcd failed, err:", err)
		return
	}
	fmt.Println("init etcd success.")

	//2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(config.LogAgent)
	//2.2 派一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现热加载配置）
	if err != nil {
		fmt.Println("get conf from etcd failed, err:", err)
		return
	}
	fmt.Println("get conf from etcd success")
	for k, v := range logEntryConf {
		fmt.Println(k, v.Topic, v.Path)
	}

	//3. 收集日志发往kafka
	//3.1 循环每一个日志收集项
	taillog.Init(logEntryConf)
	//run()
}
