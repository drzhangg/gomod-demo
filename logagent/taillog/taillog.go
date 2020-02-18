package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

func Init(fileName string) (err error) {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err :", err)
		return err
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
