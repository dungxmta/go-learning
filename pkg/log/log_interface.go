package log

import (
	"github.com/sirupsen/logrus"
	"io"
)

const (
	TRACE = logrus.TraceLevel
	DEBUG = logrus.DebugLevel
	INFO  = logrus.InfoLevel
	WARN  = logrus.WarnLevel
	ERROR = logrus.ErrorLevel
	FATAL = logrus.FatalLevel
	PANIC = logrus.PanicLevel
)

const timeFormat = "02-01-2006 15:04:05.000"

const (
	rotateMaxSize = 50 // megabytes
	rotateBackups = 10
	rotateDays    = 7
)

const EnvAssetTesting = "ENV_ASSET_TESTING"

type Config struct {
	FileOut    string      // File path; default: stdout
	Level      interface{} // default: WARN
	FormatJson bool        // Log as JSON; default: ASCII formatter

	DisableCaller      bool // Add the calling method as a field; default: always
	ShowCallerFullPath bool // Show method called with full path; default: only file name

	DisableLogFile    bool
	DisableConsoleLog bool      // os.Stdout
	IOWriter          io.Writer // add more writer; default os.Stdout & FileLog

	Rotate ConfigRotate
}

type ConfigRotate struct {
	MaxSize    int
	MaxBackups int
	MaxDays    int
	Compress   bool
	LocalTime  bool
}

type Fields logrus.Fields

type LoggerWrap interface {
	// logrus.FieldLogger
	Init(conf Config) (LoggerWrap, error)

	// implement
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields Fields) *logrus.Entry
	WithError(err error) *logrus.Entry

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}
