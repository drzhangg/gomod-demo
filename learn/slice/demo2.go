package main

import "fmt"

func main() {
	result := equal(map[string]int{"A":1}, map[string]int{"B":2})
	fmt.Println(result)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
