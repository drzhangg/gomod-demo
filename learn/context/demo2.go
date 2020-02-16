package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("db connecting....")
		time.Sleep(2 * time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker done")
}

func main() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go worker(ctx)
	time.Sleep(26 * time.Second)
	wg.Wait()
	fmt.Println("over")
}
