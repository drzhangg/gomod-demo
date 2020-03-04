package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "47.103.9.218:6379",
		Password: "",
		DB:       0,
	})
	bcmd := client.HSet("htest", "a", 123)
	fmt.Println(bcmd.Val())
	bcmd = client.HSet("htest", "b", "ccc")
	fmt.Println(bcmd.Val())

	scmd := client.HGet("htest", "b")
	fmt.Println(scmd.String())

	client.LSet()
}
