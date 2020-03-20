package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)
	//quit := make(chan bool)

	go func() {
		//写数据
		for i := 0; i < 100; i++ {
			ch <- i
			//time.Sleep(time.Second)
		}
		close(ch)
		//quit <- true
	}()

	//channel内的数据超过100时，立刻插入，不到100时，每隔2秒插入一次管道内的全部数据

	//读数据
	var clen int
	for {
		myTimer := time.NewTimer(2 * time.Second)

		clen = len(ch)
		if clen >= 10 {
			insert(ch)
		} else {
			select {
			case <-myTimer.C:
				myTimer.Reset(time.Second * 2)
				insert(ch)
			//case <-quit:
				//return
			}
		}
	}
//QUIT:
	fmt.Println("退出----")
}

func insert(ch chan int) {
	fmt.Printf("insert.....%d\n", <-ch)
}
