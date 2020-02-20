package taillog

import "gomod-demo/logagent/etcd"

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry: logEntryConf, //把当前的日志收集项配置文件保存起来
	}

	for _, logEntry := range logEntryConf {
		//logEntry.Path： 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
