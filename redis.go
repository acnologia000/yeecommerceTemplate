package main

// import (
// 	"log"
// 	"os"

// 	"github.com/go-redis/redis"
// )

// func redisConnect() *redis.Client {
// 	addr, exists := os.LookupEnv("REDIS_ADDRESS")

// 	if !exists {
// 		log.Fatal("please setup REDIS_ADDRESS environment variable")
// 	}

// 	return redis.NewClient(&redis.Options{
// 		Addr:     addr,
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})
// }

// func AddSession() string {

// 	return ""
// }
