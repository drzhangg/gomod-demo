package main

import (
	"fmt"
)

func main() {
	x := []int{1,2}
	fmt.Println(x)

	x[0],x[1] = x[1],x[0]
	fmt.Println(x)
}

