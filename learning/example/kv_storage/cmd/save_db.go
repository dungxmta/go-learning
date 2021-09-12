package cmd

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"go.uber.org/zap"

	"testProject/learning/example/kv_storage/pkg/storage"
)

const (
	defaultPath = "data_storage/badger"
)

func SaveDB(logger *zap.Logger, dbPath string, inpPath string, check bool) {
	// init store
	if dbPath == "" {
		dbPath = defaultPath
	}
	kvStore := storage.NewKvStorage(storage.Badger)
	opts := &storage.Opts{
		AutoMonitor: true,
		ReadOnly:    false,
	}
	if err := kvStore.Init(dbPath, opts); err != nil {
		logger.Fatal("Cannot init kv storage", zap.Error(err))
	}
	defer kvStore.Close()

	// open file
	fi, err := os.Open(inpPath)
	if err != nil {
		logger.Fatal("Cannot open input file", zap.Error(err), zap.String("inp_path", inpPath))
	}
	defer fi.Close()

	// save to db
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)

	var count int

	if !check { // override
		for scanner.Scan() {
			text := scanner.Text()
			logger.Debug("row", zap.String("ip", text))
			if text == "" || text == "\n" {
				continue
			}
			count += 1
			if err := kvStore.SetNull(text); err != nil {
				logger.Fatal("Something wrong when save to db", zap.Error(err), zap.Int("row_count", count))
			} else {
				logger.Info("ok", zap.Int(text, count))
			}
		}
	} else { // check exists before insert
		for scanner.Scan() {
			text := scanner.Text()
			logger.Debug("row", zap.String("ip", text))
			if text == "" || text == "\n" {
				continue
			}
			count += 1

			found, err := kvStore.Exists(text)
			if err != nil {
				logger.Fatal("Something wrong when search in db", zap.Error(err), zap.Int("row_count", count), zap.String("value", text))
			}
			if found {
				logger.Info("Added > skipping...", zap.String("value", text))
				continue
			}

			if err := kvStore.SetNull(text); err != nil {
				logger.Fatal("Something wrong when save to db", zap.Error(err), zap.Int("row_count", count))
			} else {
				logger.Info("ok", zap.Int(text, count))
			}
		}
	}

	err = scanner.Err()
	logger.Info("Save done", zap.Int("row_count", count), zap.String("inp_path", inpPath), zap.Error(err))
}

func SaveDBAndCheckMem(logger *zap.Logger, dbPath string, inpPath string, check bool) {
	PrintMemUsage()
	// init store
	if dbPath == "" {
		dbPath = defaultPath
	}
	kvStore := storage.NewKvStorage(storage.Badger)
	opts := &storage.Opts{
		AutoMonitor: true,
		ReadOnly:    false,
	}
	if err := kvStore.Init(dbPath, opts); err != nil {
		logger.Fatal("Cannot init kv storage", zap.Error(err))
	}
	defer kvStore.Close()

	// open file
	fi, err := os.Open(inpPath)
	if err != nil {
		logger.Fatal("Cannot open input file", zap.Error(err), zap.String("inp_path", inpPath))
	}
	defer fi.Close()

	// save to db
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)

	var count int

	if !check { // override
		for scanner.Scan() {
			text := scanner.Text()
			logger.Debug("row", zap.String("ip", text))
			if text == "" || text == "\n" {
				continue
			}
			count += 1
			if err := kvStore.SetNull(text); err != nil {
				logger.Fatal("Something wrong when save to db", zap.Error(err), zap.Int("row_count", count))
			} else {
				// logger.Info("ok", zap.Int(text, count))
				if count%1000 == 0 {
					logger.Info("ok", zap.Int(text, count))
					PrintMemUsage()
				}
			}
		}
	} else { // check exists before insert
		for scanner.Scan() {
			text := scanner.Text()
			logger.Debug("row", zap.String("ip", text))
			if text == "" || text == "\n" {
				continue
			}
			count += 1

			found, err := kvStore.Exists(text)
			if err != nil {
				logger.Fatal("Something wrong when search in db", zap.Error(err), zap.Int("row_count", count), zap.String("value", text))
			}
			if found {
				// logger.Info("Added > skipping...", zap.String("value", text))
				if count%1000 == 0 {
					logger.Info("Added > skipping...", zap.String("value", text))
					PrintMemUsage()
				}
				continue
			}

			if err := kvStore.SetNull(text); err != nil {
				logger.Fatal("Something wrong when save to db", zap.Error(err), zap.Int("row_count", count))
			} else {
				// logger.Info("ok", zap.Int(text, count))
				if count%1000 == 0 {
					logger.Info("ok", zap.Int(text, count))
					PrintMemUsage()
				}
			}
		}
	}

	err = scanner.Err()
	logger.Info("Save done", zap.Int("row_count", count), zap.String("inp_path", inpPath), zap.Error(err))

	PrintMemUsage()
	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
