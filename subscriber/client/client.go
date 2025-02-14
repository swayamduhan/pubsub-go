package client

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
        Password: "",
        DB:		  0, 
        Protocol: 2, 
	})

	log.Println("Connected to client at port 6379");
	testRedis(RedisClient)
}


func testRedis(RedisClient *redis.Client){
	ctx := context.Background()

	err := RedisClient.Set(ctx, "name", "abc", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisClient.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}

	log.Println("Test : ", val)

}