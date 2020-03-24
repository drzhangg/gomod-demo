package main

import "fmt"

type a struct {
	w *b
}
type b struct {
	c int
	d int
}

func main() {
	e := &a{}
	fmt.Println(e)
	fmt.Println(e.w.d)
}
