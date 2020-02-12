package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "b", "c"}
	for x, _ := range a {
		fmt.Println(x)
		//fmt.Println(y)
	}
	b := []string{"a", "b", "c","d"}
	result := equal(a, b)
	fmt.Println(result)
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
