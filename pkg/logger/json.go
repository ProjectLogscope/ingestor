package logger

import (
	"context"
	"log/slog"
)

type JSONHandler struct {
	slog.JSONHandler
}

func (h *JSONHandler) Handle(ctx context.Context, r slog.Record) error {
	if meta, ok := requestInfoFromContext(ctx); ok {
		r.AddAttrs(meta)
	}
	return h.JSONHandler.Handle(ctx, r)
}
