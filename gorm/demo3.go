package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth1 := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth1.AddDate(0, 1, -1).Format("2006-01-02")
	firstOfMonth:=firstOfMonth1.Format("2006-01-02")

	fmt.Println(lastOfMonth)
	fmt.Println(firstOfMonth)
}
