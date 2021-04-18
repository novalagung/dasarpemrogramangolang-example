package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// main function
// this describes how to connect to redis with redigo library
func main() {
	var host = "127.0.0.1"
	var port = 6379
	var username = ""
	var password = ""

	conn, err := connect(host, port, username, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Redis connected.")

	key := "key-1"
	reply, err := set(conn, key, "Hello Redis", "NX")
	if err != nil {
		fmt.Printf("redigo/set: error set value, %v", err)
		return
	}

	if reply != "OK" {
		fmt.Println("data sudah ada.")
		return
	}

	data, err := get(conn, key)
	if err != nil {
		fmt.Printf("redigo/get: error get value of key %s, %v", key, err)
		return
	}

	fmt.Println("Data of Key:", data)
}

func connect(host string, port int, username, password string) (redis.Conn, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := redis.Dial("tcp", address, redis.DialUsername(username), redis.DialPassword(password))
	if err != nil {
		return nil, fmt.Errorf("redigo/connection: error redis connection, %v", err)
	}

	return conn, nil
}

func set(conn redis.Conn, options ...interface{}) (string, error) {
	reply, err := redis.String(conn.Do("SET", options...))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func get(conn redis.Conn, key string) (string, error) {
	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}
