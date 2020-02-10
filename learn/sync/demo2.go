package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a = 0
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			lock.Lock()

			a += 1
			lock.Unlock()
			//fmt.Println(a)

			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}
	time.Sleep(time.Second)
}
