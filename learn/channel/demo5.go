package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)

	fmt.Println("now:", time.Now())

	myTricker := time.NewTicker(time.Second)//周期定时器

	i := 0

	go func() {
		for {
			nowTime := <-myTricker.C
			i++
			fmt.Println("nowTime:", nowTime)
			//quit <- true
			fmt.Println(i)

			if i == 3 {
				quit <- true//这时如果没有管道读取的话，会一直阻塞
				break
			}
		}
	}()

	<- quit
}
