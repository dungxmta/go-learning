package log

import "github.com/sirupsen/logrus"

// fields
func (ins *logger) WithField(key string, value interface{}) *logrus.Entry {
	return ins.client.WithField(key, value)
}

func (ins *logger) WithFields(fields Fields) *logrus.Entry {
	return ins.client.WithFields(logrus.Fields(fields))
}

func (ins *logger) WithError(err error) *logrus.Entry {
	return ins.client.WithError(err)
}

//  format
func (ins *logger) Debugf(format string, args ...interface{}) {
	ins.client.Debugf(format, args...)
}

func (ins *logger) Infof(format string, args ...interface{}) {
	ins.client.Infof(format, args...)
}

func (ins *logger) Printf(format string, args ...interface{}) {
	ins.client.Printf(format, args...)
}

func (ins *logger) Warnf(format string, args ...interface{}) {
	ins.client.Warnf(format, args...)
}

func (ins *logger) Warningf(format string, args ...interface{}) {
	ins.client.Warningf(format, args...)
}

func (ins *logger) Errorf(format string, args ...interface{}) {
	ins.client.Errorf(format, args...)
}

func (ins *logger) Fatalf(format string, args ...interface{}) {
	ins.client.Fatalf(format, args...)
}

func (ins *logger) Panicf(format string, args ...interface{}) {
	ins.client.Panicf(format, args...)
}

// no format
func (ins *logger) Debug(args ...interface{}) {
	ins.client.Debug(args...)
}

func (ins *logger) Info(args ...interface{}) {
	ins.client.Info(args...)
}

func (ins *logger) Print(args ...interface{}) {
	ins.client.Print(args...)
}

func (ins *logger) Warn(args ...interface{}) {
	ins.client.Warn(args...)
}

func (ins *logger) Warning(args ...interface{}) {
	ins.client.Warning(args...)
}

func (ins *logger) Error(args ...interface{}) {
	ins.client.Error(args...)
}

func (ins *logger) Fatal(args ...interface{}) {
	ins.client.Fatal(args...)
}

func (ins *logger) Panic(args ...interface{}) {
	ins.client.Panic(args...)
}

// with new line
func (ins *logger) Debugln(args ...interface{}) {
	ins.client.Debugln(args...)
}

func (ins *logger) Infoln(args ...interface{}) {
	ins.client.Infoln(args...)
}

func (ins *logger) Println(args ...interface{}) {
	ins.client.Println(args...)
}

func (ins *logger) Warnln(args ...interface{}) {
	ins.client.Warnln(args...)
}

func (ins *logger) Warningln(args ...interface{}) {
	ins.client.Warningln(args...)
}

func (ins *logger) Errorln(args ...interface{}) {
	ins.client.Errorln(args...)
}

func (ins *logger) Fatalln(args ...interface{}) {
	ins.client.Fatalln(args...)
}

func (ins *logger) Panicln(args ...interface{}) {
	ins.client.Panicln(args...)
}
