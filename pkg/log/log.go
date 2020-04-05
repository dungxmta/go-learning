package log

/* example

conf := log.Config{
	FileOut:            "info.log",
	Level:              log.INFO,
	FormatJson:         false,
	DisableLogFile:     false,
	DisableConsoleLog:  false,
	DisableRotate:      false,
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
	"io"
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

	writers := make([]io.Writer, 0)

	// set output file
	if !conf.DisableLogFile && conf.FileOut != "" {
		fPath, err := createFile(conf.FileOut)
		if err != nil {
			client.WithFields(logrus.Fields{
				"path":  fPath,
				"error": err,
			}).Error("Failed to create logs folder")
		} else {
			file, err := os.OpenFile(fPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				if conf.DisableRotate {
					// client.Out = file
					writers = append(writers, file)
				} else {
					// setup rotate
					logRotate := NewLogRotate(fPath, &conf.Rotate)
					writers = append(writers, logRotate)
				}
				client.WithField("path", fPath).Info("Config log to file ok")
			} else {
				client.WithFields(logrus.Fields{
					"path":  fPath,
					"error": err,
				}).Error("Failed to open log file, using default stdout")
			}
		}
	}

	// more writer?
	if conf.IOWriter != nil {
		writers = append(writers, conf.IOWriter)
	}

	// default disable console log
	if !conf.DisableConsoleLog {
		writers = append(writers, os.Stdout)
	}

	// set Level
	if conf.Level != nil {
		client.SetLevel(conf.Level.(logrus.Level))
	} else {
		client.SetLevel(WARN)
	}

	// enable only Stdout if env ENV_ASSET_TESTING is on
	envTesting := os.Getenv(EnvAssetTesting)

	// Output default stdout
	if len(writers) == 0 || envTesting != "" {
		client.SetOutput(os.Stdout)

		if envTesting != "" {
			client.SetLevel(DEBUG)
			client.WithField("TESTING", "on").Warn("Disable log to file for testing")
		}
	} else if len(writers) == 1 {
		// check 1 for highlight in tty (if log Stdout)
		client.SetOutput(writers[0])
	} else {
		mWriter := io.MultiWriter(writers...)
		client.SetOutput(mWriter)
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
