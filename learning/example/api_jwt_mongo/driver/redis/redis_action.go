package redis

import (
	"github.com/go-redis/redis/v7"
	"testProject/learning/example/api_jwt_mongo/driver"
)

// return new redis client -> this is diff from GetInstance() singleton
func NewClient(addr, password string, db int) (driver.MsgQueue, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// ping test
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &connector{Client: client}, nil
}

/**
error return from redis when action get nothing
i.e.
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
*/
func ErrNotFound(err error) bool {
	return err == redis.Nil
}

// implement redis action
func (ins *connector) LPush(key string, values ...interface{}) (int64, error) {
	return ins.Client.LPush(key, values).Result()
}

func (ins *connector) RPop(key string) (string, error) {
	return ins.Client.RPop(key).Result()
}

func (ins *connector) HGet(key string, field string) (string, error) {
	return ins.Client.HGet(key, field).Result()
}
