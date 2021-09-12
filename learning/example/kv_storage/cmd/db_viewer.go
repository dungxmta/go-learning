package cmd

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"

	"testProject/learning/example/kv_storage/pkg/storage"
)

func DBViewer(logger *zap.Logger, dbPath string) {
	// init store
	if dbPath == "" {
		dbPath = defaultPath
	}
	kvStore := storage.NewKvStorage(storage.Badger)
	opts := &storage.Opts{
		AutoMonitor: true,
		ReadOnly:    true,
	}
	if err := kvStore.Init(dbPath, opts); err != nil {
		logger.Fatal("Cannot init kv storage", zap.Error(err))
	}
	defer kvStore.Close()

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Type an ip to search: ")
		msg, _ := reader.ReadString('\n')

		msg = strings.Trim(msg, "")
		msg = strings.Trim(msg, "\n")
		fmt.Printf(">> Try to find ip \"%v\"\n", msg)

		found, err := kvStore.Exists(msg)
		if err != nil {
			logger.Error("Error when search", zap.Error(err))
		}
		if found {
			fmt.Println(">> ok")
		} else {
			fmt.Println(">> not found")
		}
	}

}
