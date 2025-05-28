package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(){
	RedisClient = redis.NewClient(&redis.Options{
		Addr		: os.Getenv("REDIS_ADDR"),
		Password	: os.Getenv("REDIS_PASSWORD"),
		DB		: 0,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil{
		log.Fatalf("gagal connect ke redis: %v", err)
	}

	log.Println("Berhasil connect ke redis")
}