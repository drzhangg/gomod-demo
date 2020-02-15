package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisDb *redis.Client

func main() {
	if err := initRedis(); err != nil {
		fmt.Printf("connect redis failed:%v", err)
		return
	}
	fmt.Println("connect success")

	//SetRedis()

	GetRedis()
}

func initRedis() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Password: "",
		DB:       0,
	})

	_, err = redisDb.Ping().Result()
	return
}

func SetRedis() {
	statusCmd := redisDb.Set("key", "zhang", 60*time.Second)
	fmt.Println(statusCmd.Result())
	fmt.Println(statusCmd.Name())
	fmt.Println(statusCmd.String())
	fmt.Println(statusCmd.Val())
	fmt.Println(statusCmd.Args())
}

func GetRedis() {
	stringCmd := redisDb.Get("key")
	fmt.Println(stringCmd.String())
	fmt.Println(stringCmd.Name())
	fmt.Println(stringCmd.Val())
	fmt.Println(stringCmd.Result())
}
