package main

import (
	"fmt"
	"strings"
)

func main() {
	url := `mysql://root:123456@tcp(127.0.0.1:3339)/ad_base?charset=utf8`
	splitStr := "tcp("
	s1 := strings.Split(url, splitStr)
	s2 := strings.Split(s1[1],")")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1[0]+s2[0]+s2[1])
}
