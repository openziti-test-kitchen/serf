package logger

import (
	"io"
	"log"
	"os"
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type loggerImpl struct {
	logger *log.Logger
}

func NewLogger(output io.Writer) Logger {
	return NewNamedLogger(output, "")
}

func NewNamedLogger(output io.Writer, name string) Logger {
	return NewNamedLoggerWithFlags(output, name, log.LstdFlags)
}

func NewNamedLoggerWithFlags(output io.Writer, name string, flags int) Logger {
	var logger *log.Logger
	if output != nil {
		logger = log.New(output, name, flags)
	} else {
		logger = log.New(os.Stderr, name, flags)
	}
	return &loggerImpl{logger}
}

func (l *loggerImpl) Debugf(format string, v ...interface{}) {
	l.logger.Printf("[DEBUG] serf: "+format, v...)
}

func (l *loggerImpl) Infof(format string, v ...interface{}) {
	l.logger.Printf("[INFO] serf: "+format, v...)
}

func (l *loggerImpl) Warnf(format string, v ...interface{}) {
	l.logger.Printf("[WARN] serf: "+format, v...)
}

func (l *loggerImpl) Errorf(format string, v ...interface{}) {
	l.logger.Printf("[ERROR] serf: "+format, v...)
}
