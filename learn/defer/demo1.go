package main

import "fmt"

func test1() int {
	var res int
	defer func() {
		res++
		fmt.Println(111)
	}()
	fmt.Println(res)
	return res
}

func test2() (res int) {
	defer func() {
		res++
		fmt.Println(222)
	}()
	fmt.Println(res)
	return
}

func main() {
	test1()
	test2()
}
