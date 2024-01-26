package elasticsearch

import (
	"fmt"
	"log/slog"
)

type Config struct {
	Addrs []string
}

func (c *Config) LogValue() slog.Value {
	if c == nil {
		return slog.Value{}
	}
	attrs := []slog.Attr{}
	for k, v := range c.Addrs {
		attrs = append(attrs, slog.String(fmt.Sprintf("addr.%02d", k), v))
	}
	return slog.GroupValue(attrs...)
}

type DynamicFields struct {
	keys   []string
	suffix string
}

var DefaultDynamicFields = DynamicFields{
	keys:   []string{"level", "message", "resourceId", "traceId", "spanId", "commit", "metadata_parentResourceId"},
	suffix: "_phrase",
}
