package main

import "fmt"

func main() {
	//var strChan = make(chan string)
	//go func(msg string) {
	//	strChan <- msg
	//}("hello")
	//
	//fmt.Println(<-strChan)

	//var inChan = make(chan int)
	//go loop(inChan)
	//<-inChan

	c, quit := make(chan int), make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		quit <- false
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	//<-c
	<-quit

}

func loop(in chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}

	in <- 1
}
