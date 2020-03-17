package main

import "fmt"

func main() {
	p := make([]int,2)
	//p = append(p, 1)

	p1 := p[2:]
	//p = append(p, 1)
	fmt.Println("len:", len(p))
	fmt.Println("cap:", cap(p))
	fmt.Println(p1)

	fmt.Println("len:", len(p1))
	fmt.Println("cap:", cap(p1))

	var s []string
	s = append(s, "jerry")
}
