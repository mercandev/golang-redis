package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var key = "RedisTestKey"
var value = "RedisTestValue"
var updateData = "UpdateRedisData"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		fmt.Println(err)
		return
	}

	client.Set(context.Background(), key, value, time.Minute).Err() //Set Data Redis

	val, err := client.Get(client.Context(), key).Result()

	if err != nil {
		fmt.Println(err)
		return
	}

	println(val)

	time.Sleep(10 * time.Second) //Sleep 10 Sec

	client.Set(client.Context(), key, updateData, time.Hour) //Update Key Value

	time.Sleep(10 * time.Second) //Sleep 10 Sec

	client.Del(client.Context(), key) //Delete Data
}
