package main

import (
	"fmt"
	"time"
)

var (
	x = 0
)

func read() {
	fmt.Println(x)
}

func write() {
	x += 1

}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		go write()
	}

	for i := 0; i < 1000; i++ {
		go read()
	}
	fmt.Println(time.Now().Sub(start))
}
