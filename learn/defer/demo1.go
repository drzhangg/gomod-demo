package main

import (
	"fmt"
)

func main() {
	fmt.Println("system start")
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("revocer:", msg)
		}
	}()
	pic()
}

func pic() {
	fmt.Println("1")
	panic("panic....")
	fmt.Println(2)
}
