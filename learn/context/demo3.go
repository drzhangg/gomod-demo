package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wgg sync.WaitGroup

func Worker(ctx context.Context) {
	defer wgg.Done()
	key := TraceCode("TRACE_CODE")

	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}

LOOP:
	for {
		fmt.Println("worker, trace code:", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker done")
}

func main() {
	ctx, canncel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "zhang")
	defer canncel()
	wgg.Add(1)
	defer wgg.Wait()
	go Worker(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("over")
}
