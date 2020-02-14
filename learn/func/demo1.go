package main

import "fmt"

func main() {
	f := square
	fmt.Println(f(3))
}

func square(n int) int {
	return n * n
}
