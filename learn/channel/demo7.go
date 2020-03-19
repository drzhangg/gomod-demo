package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	exit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num:", num)
			case <-time.Tick(time.Second * 3):
				exit <- true
				goto EXIT
			}
		}
	EXIT:
		fmt.Println("goto exit-------")
	}()

	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-exit

}
