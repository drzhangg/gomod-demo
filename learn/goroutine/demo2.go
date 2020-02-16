package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

var exitChan = make(chan bool, 1)

func f1() {
	defer wg1.Done()
Loop:
	for {
		fmt.Println("tom")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-exitChan:
			break Loop
		default:

		}
	}
}

func main() {
	wg1.Add(1)
	go f1()

	time.Sleep(time.Second * 5)
	exitChan <- true
	wg1.Wait()
}
