package main

import (
	"fmt"
	"sync"
)

var sg sync.WaitGroup

//任务生成器
func create_task(ch chan int) {
	sg.Wait()
	for i := 0; i < 5; i++ {
		ch <- i
	}
	//close(ch)
}

//任务执行器
func handle_task(in chan int, out chan int) {
	if len(in) != 0 {
		for temp := range in {
			out <- temp * temp
		}
	} else {
		out <- 0
	}
	//close(out)
}
func main() {
	in := make(chan int)
	out := make(chan int)

	for i := 0; i < 5; i++ {
		go create_task(in)
	}

	go handle_task(in, out)

	for r := range out {
		fmt.Println(r)
	}
}
