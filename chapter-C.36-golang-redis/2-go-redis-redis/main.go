package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// main function
// this describes how to connect to redis with redigo library
func main() {
	var host = "127.0.0.1"
	var port = 6379
	var username = ""
	var password = ""

	var ctx = context.Background()
	config := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Username: username,
		Password: password,
	}

	client := redis.NewClient(config)
	fmt.Println("Redis connected.")

	key := "key-1"
	err := client.SetNX(ctx, key, "Hello Redis", 0).Err()
	if err != nil {
		fmt.Printf("redigo/set: error set value, %v", err)
		return
	}

	data, e := client.Get(ctx, key).Result()
	if e != nil {
		fmt.Printf("redigo/get: error get value of key %s, %v", key, err)
		return
	}

	fmt.Println("Data of Key:", data)
}
