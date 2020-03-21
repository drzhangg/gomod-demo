package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cal sync.Cond

//消费者，读数据
func consumer(out <-chan int, idx int) {
	for {
		cal.L.Lock()

		for len(out) == 0 { //当缓冲区的数据取完时候，阻塞，直到有数据写入
			cal.Wait() //挂起当前协程
		}
		num := <-out
		fmt.Printf("%dth消费者，消费数据%3d，公共区域剩余%d个数据\n", idx, num, len(out))
		cal.L.Unlock()
		cal.Signal() //唤醒当前协程
	}
}

//生产者，写数据
func producer(in chan<- int, idx int) {
	for {
		cal.L.Lock()

		for len(in) == 3 { //当缓冲区的数据存满时，阻塞，等待消费者取出数据
			cal.Wait()
		}
		num := rand.Intn(1000)
		in <- num
		fmt.Printf("%dth生产者，产生数据%3d，公共区域剩余%d个数据\n", idx, num, len(in))
		cal.L.Unlock()
		cal.Signal()
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int, 3)
	quit := make(chan bool)

	cal.L = new(sync.RWMutex)

	for i := 0; i < 5; i++ {
		go producer(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		go consumer(ch, i+1)
	}

	quit <- true
}
