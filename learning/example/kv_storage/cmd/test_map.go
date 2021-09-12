package cmd

import (
	"bufio"
	"go.uber.org/zap"
	"os"
	"runtime"
)

func TestMapCheckMem(logger *zap.Logger, inpPath string) {
	PrintMemUsage()

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
	m := map[string]interface{}{}

	for scanner.Scan() {
		text := scanner.Text()
		logger.Debug("row", zap.String("ip", text))
		if text == "" || text == "\n" {
			continue
		}
		count += 1

		_, ok := m[text]
		if ok {
			// logger.Info("Added > skipping...", zap.String("value", text))
			if count%1000 == 0 {
				logger.Info("Added > skipping...", zap.String("value", text))
				PrintMemUsage()
			}
			continue
		}

		m[text] = nil
		if count%1000 == 0 {
			logger.Info("ok", zap.Int(text, count))
			PrintMemUsage()
		}
	}

	err = scanner.Err()
	logger.Info("Test done", zap.Int("row_count", count), zap.String("inp_path", inpPath), zap.Error(err))

	PrintMemUsage()
	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}
