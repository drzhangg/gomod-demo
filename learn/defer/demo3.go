package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("c")
	}()
	F()
	fmt.Println("continue")
}

func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到异常：", err)
		}
		fmt.Println("b")
	}()
	panic("a")
}
