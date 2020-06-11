package main

import "fmt"

func main() {
	in := make(chan int)
	//out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	for i := 0; i < 5; i++ {
		go func() {
			if i, ok := <-in; ok {
				fmt.Println(i)
			}
		}()
	}
}
