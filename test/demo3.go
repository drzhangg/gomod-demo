package main

import (
	"fmt"
	"time"
)

func main() {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")

	fmt.Println(year+"-"+month+"-"+day)
}
