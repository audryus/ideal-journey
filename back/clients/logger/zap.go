package logger

import (
	"fmt"

	"go.elastic.co/apm/module/apmzap/v2"
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	log, _ = zap.NewProduction(zap.WrapCore((&apmzap.Core{}).WrapCore))
}

func Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

// Info Logger
func Info(msg string, args ...interface{}) {
	sugar := log.Sugar()
	sugar.Infof(msg, args)
	log.Sync()
}

// Error logger
func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	log.Error(msg, tags...)
	log.Sync()
}
