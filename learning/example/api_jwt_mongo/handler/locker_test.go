package handler

import (
	"github.com/go-redsync/redsync/v3"
	"log"
	"sync"
	"testProject/learning/example/api_jwt_mongo/driver/redis"
	"testing"
	"time"
)

//
// refs
//  https://redis.io/topics/distlock
//  https://github.com/go-redsync/redsync/blob/master/examples/goredis/main.go
//  https://github.com/bsm/redislock
//  https://github.com/cikupin/redis-mutex-lock/blob/master/redislock/drivers/redis.go
//

func init() {
	_, err := redis.GetInstance().Init("127.0.0.1:6379", "", 1)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRedisLock(t *testing.T) {
	var wg sync.WaitGroup

	redis.GetInstance().InitLocker()

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			log.Println("===\ninit...", i)

			var opts []redsync.Option
			opts = append(opts, redsync.SetExpiry(time.Second*20))
			opts = append(opts, redsync.SetTries(5))
			opts = append(opts, redsync.SetRetryDelay(time.Second*10))

			lockKey := "lock_test"
			mux := redis.GetInstance().NewMutex(lockKey, opts...)
			err := mux.Lock()
			if err != nil {
				log.Println("-> Lock fail -> exit...", i, err)
				return
			}

			ok, err := mux.Extend()
			log.Println("extend...", i, "|", ok, "|", err)
			ok, err = mux.Valid()
			log.Println("valid...", i, "|", ok, "|", err)

			log.Println("start...", i)
			for j := 0; j < 30; j++ {
				ttl, err := redis.GetInstance().TTL(lockKey)
				log.Println("...", i, "|", ttl, "|", err)
				time.Sleep(time.Second * 1)

				if j == 10 {
					ok, err := mux.Extend()
					log.Println("extend...", i, "|", ok, "|", err)
				}
			}

			defer func() {
				ok, err := mux.Valid()
				log.Println("valid...", i, "|", ok, "|", err)

				ok, err = mux.Unlock()
				if err != nil {
					log.Println("-> Unlock fail -> exit...", i, err)
				}
				log.Println("unlock...", i, "|", ok)
			}()

		}(i)
	}

	wg.Wait()
}
