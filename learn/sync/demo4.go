package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rw sync.RWMutex

func readGo(out <-chan int, idx int) {
	for {
		rw.RLock()
		num := <-out
		fmt.Printf("-----%dth读go程，读出：%d\n", idx, num)
		rw.RUnlock()
	}
}

func writeGo(in chan<- int, idx int) {
	for {
		rw.Lock()
		//生成随机数
		num := rand.Intn(1000)
		in <- num
		fmt.Printf("%dth写go程，写入：%d\n", idx, in)
		rw.Unlock()
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go readGo(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		go writeGo(ch, i+1)
	}

	for {
		;
	}
}
