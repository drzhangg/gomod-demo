package main

import "fmt"

func main() {
	fmt.Println(f())
	fmt.Println(f1())
	fmt.Println(f2())
}

func f() (result int) {
	a := 0
	defer func() {
		result++
	}()
	return a
}

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
