package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"sync"
)

type RedisConnector struct {
	Client *redis.Client
}

var singletonRedis *RedisConnector

var once sync.Once

func GetInstance() *RedisConnector {
	once.Do(func() {
		singletonRedis = &RedisConnector{}
	})
	return singletonRedis
}

func Connect(addr, password string, db int) *RedisConnector {
	if singletonRedis == nil {
		GetInstance()
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// ping test
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect redis ok!")

	singletonRedis.Client = client
	return singletonRedis
}
