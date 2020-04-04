package main

import (
	"runtime"
	"sync"
	"testProject/pkg/log"
)

func main() {
	var wg sync.WaitGroup

	conf := log.Config{
		FileOut: "info.log",
		Level:   log.INFO,
		// FormatJson:         true,
		// DisableCaller:      true,
		// ShowCallerFullPath: true,
		// DisableLogFile:     true,
		// DisableConsoleLog:  true,
		Rotate: log.ConfigRotate{
			MaxSize:    1,
			MaxBackups: 10,
			MaxDays:    7,
			Compress:   false,
			LocalTime:  false,
		},
	}
	loggerIns, _ := log.GetInstance().Init(conf)

	runtime.GOMAXPROCS(2)

	maxNum := 1
	maxSub := 10

	wg.Add(maxNum)

	for i := 0; i < maxNum; i++ {
		go func(i int) {
			defer wg.Done()

			for j := 0; j < maxSub; j++ {
				logger := loggerIns.WithFields(log.Fields{
					"goroutine": i,
					"no":        j,
				})

				// time.Sleep(time.Second)
				logger.Info("1")
				// logger.Warn("2")
				// logger.Error("3")
			}
		}(i)
	}
	wg.Wait()
}
