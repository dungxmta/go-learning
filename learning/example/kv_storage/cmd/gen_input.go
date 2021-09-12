package cmd

import (
	"fmt"
	"go.uber.org/zap"
	"os"

	"testProject/learning/example/kv_storage/pkg/helper"
)

func GenInput(logger *zap.Logger, ip string, maxLen int) {

	output := fmt.Sprintf("input_raw/list_ips_%v.txt", maxLen)

	fo, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	if ip == "" {
		ip = "192.10.10.0"
	}

	for i := 0; i < maxLen; i++ {
		if ip == "" {
			break
		}
		fo.WriteString(fmt.Sprintf("%v\n", ip))
		logger.Info("", zap.Int(ip, i))
		ip = helper.NextIPStr(ip)
	}
}
