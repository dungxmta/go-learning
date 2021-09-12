package cmd

import (
	"go.uber.org/zap"

	"testProject/learning/example/kv_storage/pkg/storage"
)

const (
	kvPath = "learning/example/kv_storage/data_storage/badger"
)

func Run(logger *zap.Logger) {

	kvStore := storage.NewKvStorage(storage.Badger)
	opts := &storage.Opts{
		AutoMonitor: true,
		ReadOnly:    false,
	}
	if err := kvStore.Init(kvPath, opts); err != nil {
		logger.Fatal("Cannot init kv storage", zap.Error(err))
	}
	defer kvStore.Close()

	// TODO
	// 1. mem test using map
	// 2. mem test using kv storage
}
