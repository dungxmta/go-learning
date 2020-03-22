package redis

import (
	"github.com/go-redis/redis/v7"
	"sync"
	"testProject/learning/example/api_jwt_mongo/driver"
)

type RedisConnector struct {
	Client *redis.Client
}

var singleton driver.MsgQueue

var once sync.Once

func GetInstance() driver.MsgQueue {
	once.Do(func() {
		singleton = &RedisConnector{}
	})
	return singleton
}

func (ins *RedisConnector) Init(addr, password string, db int) (driver.MsgQueue, error) {
	ins.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// ping test
	_, err := ins.Client.Ping().Result()
	return ins, err
}
