package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x    = 0
	wg   sync.WaitGroup
	lock sync.Mutex
)

func read() {
	fmt.Println(x)
	defer wg.Done()
}

func write() {
	defer wg.Done()
	lock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}

	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
