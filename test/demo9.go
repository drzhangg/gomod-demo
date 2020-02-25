package main

import (
	"fmt"
	"strings"
)

func main() {
	arr1 := []string{"jerry", "tom", "bob"}
	str1 := strings.Join(arr1, ",")
	fmt.Println(str1)
}
