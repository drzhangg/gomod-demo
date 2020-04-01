package main

import "fmt"

func main() {
	a := []int{1}

	a = append(a, 2)
	a = append(a, 3)

	b := append(a,4)
	fmt.Printf("%p\n",a)
	fmt.Printf("%p\n",b)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
}
