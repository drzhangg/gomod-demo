package main

import (
	"fmt"
	"time"
)

func main() {
	//传输数据channel
	ch := make(chan int)
	//退出channel
	quit := make(chan bool)

	//子go程，写数据
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true //向退出程序的channel里面写入数据，当主go程读取到数据的时候，退出程序
	}()

	//主go程，读数据

	for {
		select {
		case num := <-ch:
			fmt.Println("读到：", num)
		case <-quit:
			goto EXIT
		}
		fmt.Println("========")
	}
EXIT:
}
