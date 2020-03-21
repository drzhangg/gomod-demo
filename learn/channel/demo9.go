package main

import (
	"fmt"
	"time"
)

var ch  = make(chan int)

func printer(s string) {
	for _, v := range s {
		fmt.Printf("%c", v)
		time.Sleep(time.Millisecond * 300)
	}

}

func person1() {
	printer("hello")
	ch <- 66
}

func person2() {
	<- ch
	printer("world")
}

func main() {
	go person1()
	go person2()
	for {
		;
	}
}
