package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var flag bool

func f() {
	defer wg.Done()

	for {
		fmt.Println("jerry")
		time.Sleep(500 * time.Millisecond)

		if flag {
			break
		}
	}

}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(5 * time.Second)

	flag = true
	wg.Wait()
}
