package taillog

import (
	"fmt"
	"gomod-demo/logagent/etcd"
	"time"
)

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, //把当前的日志收集项配置文件保存起来
		tskMap: make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}

	for _, logEntry := range logEntryConf {
		//logEntry.Path： 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}

	go tskMgr.run()

}

// 通过一直监听自己的newConfChan，有了新的配置过来之后做相应的处理
func (t *tailLogMgr) run() {
	//因为要一直监听是否有配置发送过来，所以要用个for
	for {
		select {
		case newConf := <-t.newConfChan: //当有新的数据过来时，会通过管道传递
			//1.配置新增
			//2.配置删除
			//3.配置变更
			fmt.Println("新的配置来了！", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数，向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
