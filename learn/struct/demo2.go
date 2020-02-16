package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*Student)

	stus := []*Student{{"jerry", 12}, {"tom", 13}, {"zhang", 23}}

	for _, stu := range stus {
		m[stu.name] = stu
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
