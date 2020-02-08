package main

import (
	"fmt"
	"gomod-demo/demo"
)

func main() {
	demo.Service()

	v1 := make([]int,0)
	fmt.Println(len(v1))
	fmt.Println(cap(v1))

	fmt.Println("---------")

	v1 = append(v1, 1)
	fmt.Println(len(v1))
	fmt.Println(cap(v1))
	fmt.Println("---------")

	v1 = append(v1, 2)
	fmt.Println(len(v1))
	fmt.Println(cap(v1))
	fmt.Println("---------")
	v1 = append(v1, 3)
	fmt.Println(len(v1))
	fmt.Println(cap(v1))

}
