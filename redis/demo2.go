package main

import (
	"github.com/go-redis/redis"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})

	//client.
	client.Set("name", "jerry", time.Second * 5)
	//client.Del("name")
}
