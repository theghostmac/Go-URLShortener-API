package storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Storage struct {
	redisClient *redis.Client
}

var (
	store = &Storage{}
	contx = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitStore() *Storage {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:8080",
		Password: " ",
		DB:       0,
	})
	pinger, err := redisClient.Ping(contx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing Redis: %v", err))
	}
	fmt.Printf("\nRedis initialized successfully. Pinging {%s}", pinger)
	store.redisClient = redisClient
	return store
}

func StoreUrl(shortURL string, originalURL string, userID string) {
	err := store.redisClient.Set(contx, shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Url key not saved. Error: %v - shortURL: %s - originalURL: %s\n", err, shortURL, originalURL))
	}
}

func ReceiveUrl(shortURL string) string {
	result, err := store.redisClient.Get(contx, shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Url key  not retrieved. Error: %v - shortURL: %s", err, shortURL))
	}
	return result
}
