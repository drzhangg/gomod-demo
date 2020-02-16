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

	key := "rank"
	items := []redis.Z{
		redis.Z{
			Score:  99,
			Member: "golang",
		},
		redis.Z{
			Score:  90,
			Member: "java",
		},
		redis.Z{
			Score:  80,
			Member: "php",
		},
	}

	redisDb.ZAdd(key, items...)

	newScore, err := redisDb.ZIncrBy(key, 10.0, "golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed:%s", err)
		return
	}
	fmt.Println("golang score is: ", newScore)
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
	fmt.Println(stringCmd.Val())
}
