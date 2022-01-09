package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func newRedisClient(redisHost string, redisPassword string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	return client
}

func main() {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	rdb := newRedisClient(redisHost, redisPassword)
	fmt.Println("redis client initialized")

	key := "key-1"
	data := "Hello Redis"
	ttl := time.Duration(3) * time.Second

	// store data using SET command
	op1 := rdb.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")

	// get data
	op2 := rdb.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	res, err := op2.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	log.Println("get operation success. result:", res)
}
