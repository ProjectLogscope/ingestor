package config

import (
	"fmt"
	"strings"
)

const (
	LogLevelDebug string = "debug"
	LogLevelWarn  string = "warn"
	LogLevelInfo  string = "info"
	LogLevelError string = "error"
)

type ServiceLogLevel struct {
	level string
}

func (s *ServiceLogLevel) UnmarshalText(b []byte) error {
	level := strings.ToLower(string(b))
	switch level {
	case
		LogLevelDebug, LogLevelWarn, LogLevelInfo, LogLevelError:
		s.level = level
		return nil
	default:
		return fmt.Errorf("error processing environment variable SERVER_LOG_LEVEL: invalid or missing value %q", level)
	}
}

func (s *ServiceLogLevel) GetLevel() string {
	return s.level
}
