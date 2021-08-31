package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	request := "finno-identity-image-service:save-photo:4436:400:12efa4bced1397bf"

	test := "*4436:400*"

	if err := redisClient.Set(request, nil, 10*time.Minute).Err(); err != nil {
		fmt.Println("set key redis error : ", err)
	}

	exist, err := redisClient.Keys(test).Result()
	if err != nil {
		fmt.Println("check exist redis error")
	}

	if len(exist) <= 3 {
		fmt.Println("Fail")
	} else {
		fmt.Println("Success")
	}

	// s, _ := json.MarshalIndent(m, "", "\t")
	// fmt.Println(string(s))
}
