package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))


	 var a [10]int
	fmt.Println(len(a))
	fmt.Println(a)
}