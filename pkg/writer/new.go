package writer

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func New(logfile string) io.Writer {
	return &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    LoggerMaxSize,
		MaxBackups: LoggerMaxBackups,
		MaxAge:     LoggerMaxAge,
		Compress:   LoggerCompress,
	}
}
