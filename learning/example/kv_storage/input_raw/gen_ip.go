package main

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"

	"testProject/learning/example/kv_storage/pkg/helper"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	now := time.Now()

	defer func() {
		d := time.Since(now)
		logger.Info("END", zap.Duration("time", d))
	}()

	output := "list_ips.txt"

	fo, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	ip := "192.10.10.0"

	for i := 0; i < 1000000; i++ {
		if ip == "" {
			break
		}
		fo.WriteString(fmt.Sprintf("%v\n", ip))
		logger.Info("", zap.Int(ip, i))
		ip = helper.NextIPStr(ip)
	}
}
