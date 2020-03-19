package main

import "fmt"

// 发送数据，将数据放到channel里面
func send(out chan<- int) {
	out <- 12
}

//接收数据，从channel里面读取数据
func receive(in <-chan int) {
	num := <-in
	fmt.Println(num)
}

func main() {
	ch := make(chan int)

	go func() {
		send(ch)
	}()

	receive(ch)
}
