package main

import "fmt"

type OrderInfo struct {
	id int
}

//生产者  将数据写入到channel
func producer(in chan<- OrderInfo) {
	for i := 0; i < 10; i++ {
		in <- OrderInfo{id: i + 1}
	}

	close(in)
}

//消费者   从channel中读取数据
func consumer(out <-chan OrderInfo) {

	for {
		if num, ok := <-out; ok {
			fmt.Println("订单id为：", num.id)
		}else {
			n := <- out
			fmt.Println("结束：",n)
			break
		}
	}
	//for order := range out {
	//	fmt.Println("订单id为：", order.id)
	//}
}

func main() {
	order := make(chan OrderInfo, 10)

	go producer(order)

	consumer(order)
}
