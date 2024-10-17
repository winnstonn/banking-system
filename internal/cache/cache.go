package cache

import (
	"log"

	"github.com/redis/go-redis/v9"
)

type ICache interface {
	RetrieveData() string
	Close()
}

type Cache struct {
	redisConn *redis.Client
}

func NewCache() ICache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Cache{
		redisConn: redisClient,
	}
}

func (c *Cache) Close() {
	err := c.redisConn.Close()
	if err != nil {
		log.Println("Failed to close redis connection", err)
		return
	}
	log.Println("shutting Redis down")
}
