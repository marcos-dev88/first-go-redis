package main

import (
	"log"
	"time"

	"github.com/marcos-dev88/first-go-redis/cache"
)

var redisURI = "redis-cache-poc-go:6079"

func main() {

	redisCall := cache.NewRedis(0, redisURI, "somePass")

	c := cache.NewCache(redisCall)

	err := c.Create("myaddr6", "something_6", 0)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	time.Sleep(10 * time.Minute)
}
