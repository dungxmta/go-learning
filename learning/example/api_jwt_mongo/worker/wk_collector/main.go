package main

import (
	"fmt"
	// "log"
	"testProject/learning/example/api_jwt_mongo/config"
	driverRedis "testProject/learning/example/api_jwt_mongo/driver/redis"
	"time"
	// "context"
)

const (
	QUEUE = "plugin_config"
)

// TODO: https://github.com/robfig/cron

func main() {
	redisAddr := config.GetStr("REDIS_ADDR")
	redisPwd := config.GetStr("REDIS_PWD")
	redisDB := int(config.Get("REDIS_DB").(float64))

	fmt.Println(redisAddr)

	redis := driverRedis.Connect(redisAddr, redisPwd, redisDB)

	count := 1
	// map table with plugin name & instance
	pluginMap := make(map[string]Job)

	for {
		pluginConf, _ := redis.Client.RPop(QUEUE).Result()
		// TODO: "redis: nil" ???
		// if err.Error() != "redis: nil" {
		// 	fmt.Println(err)
		// }

		if pluginConf == "" {
			fmt.Println("...main...")
			time.Sleep(time.Second * 5)
			continue
		}

		// TODO: check config type here
		// plugin per goroutine
		id := fmt.Sprintf("plugin_%v", count)
		fmt.Println(id, " | ", pluginConf)

		plugin := NewPlugin(id, pluginConf)

		go plugin.Process()

		time.AfterFunc(time.Second*5, func() {
			pluginMap[id].Stop()
		})

		pluginMap[id] = plugin

		count++
	}
}

type Job interface {
	Process()
	Stop()
}

type Plugin struct {
	id   string
	conf string
	stop chan bool
}

func NewPlugin(id string, conf string) *Plugin {
	return &Plugin{
		id:   id,
		conf: conf,
		stop: make(chan bool),
	}
}

func (p *Plugin) Process() {
	defer fmt.Println("-> End", p.id)
	fmt.Println("-> Start", p.id)
	fmt.Println("...running ", p.id, "with config", p.conf)

	timeout := time.After(23 * time.Hour)
	// TODO: logic here
	// for i := 0; i < 20; i++ {
	for {
		select {
		case <-p.stop: // wait stop/done signal
			fmt.Println("+ Plugin", p.id, "get Stop signal!")
			return
		case <-timeout:
			fmt.Println("out of time :(")
		default:
			// TODO: logic here
			time.Sleep(time.Second)
			fmt.Println("... <-", p.id)
		}
	}
}

func (p *Plugin) Stop() {
	defer fmt.Println("Done stopping", p.id)
	fmt.Println("... stopping", p.id)
	p.stop <- true
}

// func (p *Plugin) Process() {
// 	defer fmt.Println("-> End", p.id)
// 	fmt.Println("-> Start", p.id)
//
// 	// ***NOTE: cancel parent context not cancel its child goroutine
// 	// cancel context not mean return func, only send value to channel context.Done()
// 	ctx, cancel := context.WithCancel(context.Background())
//
// 	// time.AfterFunc(time.Second*2, func() {
// 	// 	log.Println("Cancel parent context after 2s")
// 	// 	cancel()
// 	// })
//
// 	// wait stop/done signal
// 	go func(ctx context.Context, cancelFunc context.CancelFunc, p *Plugin) {
// 		for {
// 			select {
// 			case stopSignal := <-p.stop:
// 				fmt.Println("Plugin", p.id, "get Stop signal!", stopSignal)
// 				cancelFunc()
// 				return
// 			case <-ctx.Done():
// 				fmt.Println("Plugin", p.id, "get Done signal!")
// 				return
// 			default:
// 				fmt.Println("... <-", p.id, "wait signal...")
// 				time.Sleep(time.Second)
// 			}
// 		}
// 	}(ctx, cancel, p)
//
// 	// run logic of plugin
// 	// when logic done, callback cancel to exit parent
// 	// go func(p *Plugin, cancelFunc context.CancelFunc, c context.Context) {
// 	// 	defer cancelFunc() // cancelFunc() --> parent context.Done() --> exit parent
//
// 	// TODO: logic here
// 	// logic(ctx, cancel, p)
// 	fmt.Println("...running ", p.id, "with config", p.conf)
// 	for i := 0; i < 20; i++ {
// 		select {
// 		case <- ctx.Done():
// 			fmt.Println("+ Plugin", p.id, "get Done signal!")
// 			return
// 		case <-time.After(23 * time.Hour):
// 			fmt.Println("out of time :(")
// 		default:
// 			time.Sleep(time.Second)
// 			fmt.Println("... <-", p.id)
// 		}
// 	}
// 	// }(p, cancel, ctxChild)
// 	return
// }
