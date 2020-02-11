package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	for x,_ := range a {
		fmt.Println(x)
		//fmt.Println(y)
	}
}
