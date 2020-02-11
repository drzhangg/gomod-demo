package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 2)
	var lock sync.Mutex

	go func() {
		lock.Lock()
		fmt.Println("goroutine1: 我会锁定大概 2s")
		time.Sleep(2 * time.Second)
		fmt.Println("我解锁了，你们去抢吧")
		defer lock.Unlock()
		ch <- struct{}{}
	}()

	go func() {
		fmt.Println("goroutine2：等待1解锁")
		lock.Lock()
		defer lock.Lock()
		fmt.Println("goroutine2:我也解锁了")
		ch <- struct{}{}
	}()

	for i := 0; i < 2; i++ {
		<-ch
	}
}
