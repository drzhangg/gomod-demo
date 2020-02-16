package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

func worker(ctx context.Context) {
	defer wg3.Done()
Loop:
	for {
		fmt.Println("bob")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			break Loop
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	wg3.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg3.Wait()
}
