package redis

import (
	redisLib "github.com/go-redis/redis"
	"github.com/go-redsync/redsync/v3"
	"github.com/go-redsync/redsync/v3/redis"
	"github.com/go-redsync/redsync/v3/redis/goredis"
	"sync"
	"testProject/learning/example/api_jwt_mongo/driver"
)

type connector struct {
	Client *redisLib.Client
	Locker *redsync.Redsync
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
	ins.Client = redisLib.NewClient(&redisLib.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// ping test
	_, err := ins.Client.Ping().Result()
	return ins, err
}

func (ins *connector) InitLocker() {
	pool := goredis.NewGoredisPool(ins.Client)
	ins.Locker = redsync.New([]redis.Pool{pool})
}

func (ins *connector) NewMutex(name string, options ...redsync.Option) driver.Locker {
	return ins.Locker.NewMutex(name, options...)
}
