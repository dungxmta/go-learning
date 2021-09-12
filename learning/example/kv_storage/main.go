package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"

	lstCmd "testProject/learning/example/kv_storage/cmd"
)

func main() {
	// TODO
	// 1. gen storage from input raw
	// 2. mem test using map
	// 3. mem test using kv storage
	// 4. speed test using kv storage

	// logger, _ := zap.NewProduction()
	var logCfg = zap.NewProductionConfig()
	logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, _ := logCfg.Build()
	defer logger.Sync() // flushes buffer, if any

	now := time.Now()

	defer func() {
		d := time.Since(now)
		logger.Info("END", zap.Duration("time", d))
	}()

	// command gen IP
	var (
		fromIP string
		maxLen int
	)

	var cmdGenInput = &cobra.Command{
		Use:   "gen_input",
		Short: "Generate IP to file",
		// Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lstCmd.GenInput(logger, fromIP, maxLen)
		},
	}
	cmdGenInput.Flags().StringVarP(&fromIP, "from_ip", "f", "192.10.10.0", "generate IP")
	cmdGenInput.Flags().IntVarP(&maxLen, "max_len", "l", 10, "Number of IPs to generate")

	// command gen IP
	var (
		dbPath  string
		inpPath string
		check   bool
		showMem bool
	)

	var cmdSaveDB = &cobra.Command{
		Use:   "save_db",
		Short: "Save input data to DB",
		// Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if showMem {
				lstCmd.SaveDBAndCheckMem(logger, dbPath, inpPath, check)
				return
			}
			lstCmd.SaveDB(logger, dbPath, inpPath, check)
		},
	}
	cmdSaveDB.Flags().StringVarP(&dbPath, "db_path", "d", "data_storage/badger", "DB path")
	cmdSaveDB.Flags().StringVarP(&inpPath, "inp_path", "i", "input_raw/list_ips_5.txt", "Input file path")
	cmdSaveDB.Flags().BoolVarP(&check, "check", "c", false, "Check existed each value before add")
	cmdSaveDB.Flags().BoolVarP(&showMem, "mem", "m", false, "Show memory usage")

	// command viewer
	var cmdDBViewer = &cobra.Command{
		Use:   "db_viewer",
		Short: "DB viewer",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			lstCmd.DBViewer(logger, dbPath)
		},
	}
	cmdDBViewer.Flags().StringVarP(&dbPath, "db_path", "d", "data_storage/badger", "DB path")

	// command run
	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "Run test",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			lstCmd.Run(logger)
		},
	}

	// default command
	rootCmd := &cobra.Command{Use: ""}
	rootCmd.AddCommand(
		cmdRun,
		cmdGenInput,
		cmdSaveDB,
		cmdDBViewer,
	)

	rootCmd.Execute()
}
