package main

import "fmt"

func main() {
	p := make([]int,1)
	//p = append(p, 1)
	fmt.Println("len:", len(p))
	fmt.Println("cap:", cap(p))
	fmt.Println(p[1:])

}
