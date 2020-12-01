package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for _ = range ticker.C {
			select {
			case d := <-done:
				fmt.Println("done:", d)
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message:%d \n", <-messages)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		messages <- i
	}

	time.Sleep(5 * time.Second)
	close(done)

	time.Sleep(time.Second)
	fmt.Println("main process exit")

}
