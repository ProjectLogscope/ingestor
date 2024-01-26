package logger

import (
	"io"
	"log/slog"
)

func New(writer io.Writer, ht handlerType, sopts *slog.HandlerOptions) slog.Handler {
	var handler slog.Handler
	switch ht {
	case TypeText:
		handler = &TextHandler{
			TextHandler: *slog.NewTextHandler(writer, sopts),
		}
	case TypeJSON:
		handler = &JSONHandler{
			JSONHandler: *slog.NewJSONHandler(writer, sopts),
		}
	}
	return handler
}
