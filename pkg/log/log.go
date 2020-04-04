package log

/* example

conf := log.Config{
	FileOut:      "info.log",
	Level:        log.INFO,
	FormatJson:   false,
}
logger, _ := log.GetInstance().Init(conf)

logger.Info("1")
logger.Infof("%v", "2)
logger.Infoln("3)

loggerWithFields := logger.WithFields(log.Fields{
	"common_field_01": 1,
	"common_field_02": 2,
})
loggerWithFields.Info("This log will have 2 common fields")

*/

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

// object
type logger struct {
	client logrus.FieldLogger
	// sync.Mutex
}

var singleton LoggerWrap
var once sync.Once

// get/set
func GetInstance() LoggerWrap {
	once.Do(func() {
		singleton = NewLogger()
	})
	return singleton
}

func SetInstance(ins LoggerWrap) {
	singleton = ins
}

// new object
func NewLogger() LoggerWrap {
	return &logger{client: logrus.New()}
}

// init instance
func (ins *logger) Init(conf Config) (LoggerWrap, error) {
	// new logger
	client := logrus.New()

	// set output file
	var confToFile bool
	if conf.FileOut != "" {
		fpath, _ := createFile(conf.FileOut)
		file, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			client.Out = file
			confToFile = true
		} else {
			client.Warn("Failed to log to file, using default stdout")
		}
	}

	// Fallback to IOWriter if confToFile fail
	if !confToFile {
		// Output to stdout instead of the default stderr
		// Can be any io.Writer
		if conf.IOWriter != nil {
			client.SetOutput(conf.IOWriter)
		} else {
			client.SetOutput(os.Stdout)
		}
	}

	// set Level
	if conf.Level != nil {
		client.SetLevel(conf.Level.(logrus.Level))
	} else {
		client.SetLevel(logrus.WarnLevel)
	}

	// Add the calling method as a field
	client.SetReportCaller(!conf.DisableCaller)

	// Log as JSON or the default ASCII formatter
	var formatter logrus.Formatter

	if conf.FormatJson {
		formatter = &logrus.JSONFormatter{
			TimestampFormat:   timeFormat,
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			FieldMap:          nil,
			PrettyPrint:       false,
		}
		// split method called to only name instead of full path
		if !conf.DisableCaller && !conf.ShowCallerFullPath {
			formatter.(*logrus.JSONFormatter).CallerPrettyfier = callerOnlyFileName
		}
		client.SetFormatter(formatter)
	} else {
		formatter = &logrus.TextFormatter{
			TimestampFormat:  timeFormat,
			DisableTimestamp: false,
		}
		if !conf.DisableCaller && !conf.ShowCallerFullPath {
			formatter.(*logrus.TextFormatter).CallerPrettyfier = callerOnlyFileName
		}
	}
	client.SetFormatter(formatter)

	ins.client = client
	return ins, nil
}
