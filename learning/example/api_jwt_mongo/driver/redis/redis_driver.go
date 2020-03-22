package redis

import (
	"github.com/go-redis/redis/v7"
	"sync"
	"testProject/learning/example/api_jwt_mongo/driver"
)

type connector struct {
	Client *redis.Client
}

var singleton driver.MsgQueue

var once sync.Once

func GetInstance() driver.MsgQueue {
	once.Do(func() {
		singleton = &connector{}
	})
	return singleton
}

func (ins *connector) Init(addr, password string, db int) (driver.MsgQueue, error) {
	ins.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// ping test
	_, err := ins.Client.Ping().Result()
	return ins, err
}
