package main

import "fmt"

func rescuive(i int) int {
	if i <= 1 {
		return 1
	}
	return rescuive(i-1) + rescuive(i-2)
}

func main() {
	for i := 0; i < 6; i++ {
		sum := rescuive(i)
		fmt.Println(sum)
	}
}
