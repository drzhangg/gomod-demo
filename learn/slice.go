package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	var z []int
	zlen := len(a) + 1
	z = a[:zlen]
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(z)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
