package log

import "github.com/golang/glog"

var Log Logger = &glogLogger{}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type glogLogger struct {
}

func (l *glogLogger) Debug(args ...interface{}) {
	panic("unimplemented")
}

func (l *glogLogger) Info(args ...interface{}) {
	glog.Info(args)
}

func (l *glogLogger) Warning(args ...interface{}) {
	glog.Warning(args)
}

func (l *glogLogger) Error(args ...interface{}) {
	glog.Error(args)
}

func (l *glogLogger) Fatal(args ...interface{}) {
	glog.Fatal(args)
}
