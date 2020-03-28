package main

import "fmt"

var b = 100

const c = 123

func main() {
	a := make(map[string]string)
	fmt.Println(&b, b)
	//fmt.Println(&c, c)
	fmt.Println(&a, a)
}
