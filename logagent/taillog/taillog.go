package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"gomod-demo/logagent/kafka"
)

type TailTask struct {
	path, topic string
	instance    *tail.Tail
}

var (
	tailObj *tail.Tail
	LogChan chan string
)

// 构造函数
func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init()
	return tailObj
}

func (t *TailTask) init() {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
	}

	go t.run() //直接去采集日志发送到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines: //从tailObj的通道中一行一行的读取日志数据
			// kafka.SentToKafka(t.topic, line.Text) 函数调用函数
			// 先把日志数据发送到一个通道中
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
