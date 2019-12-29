package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Demo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	N = "a_%s"
)

func main() {
	a := "1"

	N = fmt.Sprintf(N,a)
	fmt.Println(N)
}
